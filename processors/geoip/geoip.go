package geoip

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/blushft/strana"
	"github.com/blushft/strana/domain/entity"
	"github.com/blushft/strana/platform/config"
	"github.com/blushft/strana/processors"
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
	processors.Register("geoip", new)
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
	opts Options
	db   *geo.Reader
}

func new(conf config.Processor) (strana.Processor, error) {
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
	}, nil
}

func (p *geoproc) Process(msg *entity.RawMessage) ([]*entity.RawMessage, error) {
	if msg == nil {
		return nil, nil
	}

	if msg.IPAddress == "" {
		return []*entity.RawMessage{msg}, nil
	}

	ip := net.ParseIP(msg.IPAddress)

	city, err := p.db.City(ip)
	if err != nil {
		return nil, err
	}

	msg.City = city.City.Names[p.opts.Language]
	msg.Country = city.Country.Names[p.opts.Language]
	msg.Region = city.Continent.Names[p.opts.Language]
	msg.PostalCode = city.Postal.Code
	msg.Timezone = city.Location.TimeZone
	msg.Latitude = strconv.FormatFloat(city.Location.Latitude, 'f', -1, 64)
	msg.Longitude = strconv.FormatFloat(city.Location.Longitude, 'f', -1, 64)

	return []*entity.RawMessage{msg}, nil
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
