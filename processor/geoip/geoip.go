package geoip

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/blushft/strana/event"
	"github.com/blushft/strana/event/contexts"
	"github.com/blushft/strana/platform"
	"github.com/blushft/strana/platform/config"
	"github.com/blushft/strana/processor"
	"github.com/mitchellh/mapstructure"

	geo "github.com/oschwald/geoip2-golang"
	"github.com/pkg/errors"
)

const (
	mmdbPermalink       = "https://download.maxmind.com/app/geoip_download?edition_id=%s&license_key=%s&suffix=tar.gz"
	mmdbGeoLiteCity     = "GeoLite2-City"
	mmdbGeoLiteCityFile = "GeoLite2-City.mmdb"
)

func init() {
	platform.RegisterEventProcessor("geoip", new)
}

type Options struct {
	MaxMindLicense    string `json:"max_mind_license" yaml:"max_mind_license" mapstructure:"max_mind_license" yaml.mapstructure:"max_mind_license"`
	DatabasePath      string `json:"database_path" yaml:"database_path" mapstructure:"database_path" yaml.mapstructure:"database_path"`
	AutomaticDownload bool   `json:"automatic_download" yaml:"automatic_download" mapstructure:"automatic_download" yaml.mapstructure:"automatic_download"`
	Language          string `json:"language" yaml.mapstructure:"language"`
}

func newOptions(m map[string]interface{}) Options {
	opts := Options{
		DatabasePath:      "./",
		AutomaticDownload: false,
		Language:          "en",
	}

	var copts Options
	if err := mapstructure.Decode(m, &copts); err != nil {
		return opts
	}

	if copts.MaxMindLicense != "" {
		copts.AutomaticDownload = true
	}

	if copts.DatabasePath == "" {
		copts.DatabasePath = opts.DatabasePath
	}

	if copts.Language == "" {
		copts.Language = opts.Language
	}

	return copts
}

type geoproc struct {
	opts      Options
	db        *geo.Reader
	validator event.Validator
}

func new(conf config.Processor) (processor.EventProcessor, error) {
	opts := newOptions(conf.Options)

	dbpath := path.Join(opts.DatabasePath, mmdbGeoLiteCityFile)
	if !exists(dbpath) {
		if opts.AutomaticDownload {
			if err := getMMDB(opts); err != nil {
				return nil, err
			}
		} else {
			return nil, errors.New("unable to open geolite2 database: invalid path")
		}
	}

	db, err := geo.Open(dbpath)
	if err != nil {
		return nil, err
	}

	return &geoproc{
		opts: opts,
		db:   db,
		validator: event.NewValidator(
			event.WithRules(map[string]event.Rule{
				"has_network": event.HasContext(contexts.ContextNetwork),
				"has_ip":      event.ContextContains(contexts.ContextNetwork, "ip", true),
			}),
		),
	}, nil
}

func (p *geoproc) Process(evt *event.Event) ([]*event.Event, error) {
	if evt == nil {
		return nil, nil
	}

	if !p.validator.Validate(evt) {
		return []*event.Event{evt}, nil
	}

	v := evt.Context["network"].Interface()
	netctx := v.(*contexts.Network)

	city, err := p.db.City(netctx.IP)
	if err != nil {
		return nil, err
	}

	st := ""
	if len(city.Subdivisions) > 0 {
		st = city.Subdivisions[0].IsoCode
	}

	locctx := &contexts.Location{
		City:       city.City.Names[p.opts.Language],
		State:      st,
		Country:    city.Country.Names[p.opts.Language],
		Region:     city.Continent.Names[p.opts.Language],
		Locale:     strconv.Itoa(int(city.Location.MetroCode)),
		PostalCode: city.Postal.Code,
		Timezone:   city.Location.TimeZone,
		Latitude:   city.Location.Latitude,
		Longitude:  city.Location.Longitude,
	}

	evt.SetContext(locctx)

	return []*event.Event{evt}, nil
}

func getMMDB(opts Options) error {
	url := fmt.Sprintf(mmdbPermalink, mmdbGeoLiteCity, opts.MaxMindLicense)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	outpath := path.Join(os.TempDir(), "mmdb_geolite_dl")

	out, err := os.Create(outpath)
	if err != nil {
		return err
	}

	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return extract(outpath, opts)
}

func extract(archivePath string, opts Options) error {
	gzipStr, err := os.Open(archivePath)
	if err != nil {
		return err
	}

	uncompressedStream, err := gzip.NewReader(gzipStr)
	if err != nil {
		return err
	}

	tarReader := tar.NewReader(uncompressedStream)

	if !exists(opts.DatabasePath) {
		if err := os.MkdirAll(opts.DatabasePath, 0755); err != nil {
			return err
		}
	}

	var extractDir string

	for true {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		switch header.Typeflag {
		case tar.TypeDir:
			extractDir = path.Join(archivePath, header.Name)
			if err := os.MkdirAll(extractDir, 0755); err != nil {
				return err
			}
		case tar.TypeReg:
			fileout := path.Join(archivePath, header.Name)
			outFile, err := os.Create(fileout)
			if err != nil {
				return err
			}
			if _, err := io.Copy(outFile, tarReader); err != nil {
				return err
			}
			outFile.Close()

		default:
			return errors.Errorf(
				"uknown type: %s in %s",
				string(header.Typeflag),
				header.Name)
		}
	}

	dbfile := path.Join(extractDir, mmdbGeoLiteCityFile)
	if !exists(dbfile) {
		return errors.New("could not locate geolite file after extraction: " + dbfile)
	}

	mvpath := path.Join(opts.DatabasePath, mmdbGeoLiteCityFile)
	if err := moveFile(dbfile, mvpath); err != nil {
		return errors.Wrapf(err, "unable to move %s to %s", dbfile, mvpath)
	}

	return nil
}

func exists(p string) bool {
	_, err := os.Stat(p)
	return !os.IsNotExist(err)
}

func moveFile(src, dst string) error {
	input, err := os.Open(src)
	if err != nil {
		return errors.Wrap(err, "move file open source")
	}

	output, err := os.Create(dst)
	if err != nil {
		input.Close()
		return err
	}

	defer func() {
		input.Close()
		output.Close()
	}()

	_, err = io.Copy(output, input)
	if err != nil {
		return errors.Wrap(err, "move file io copy")
	}

	return nil
}
