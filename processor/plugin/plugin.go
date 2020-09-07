package plugin

import (
	"io/ioutil"
	"path/filepath"
	"reflect"

	"github.com/blushft/strana/event"
	"github.com/blushft/strana/platform/config"
	"github.com/blushft/strana/processor"
	"github.com/blushft/strana/processor/plugin/imports"
	"github.com/containous/yaegi/interp"
	"github.com/containous/yaegi/stdlib"
	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
)

type Manifest struct {
	Name        string                 `json:"name" structs:"name" mapstructure:"name"`
	Description string                 `json:"description" structs:"description" mapstructure:"description"`
	Version     string                 `json:"version" structs:"version" mapstructure:"version"`
	Type        string                 `json:"type" structs:"type" mapstructure:"type"`
	Import      string                 `json:"import" structs:"import" mapstructure:"import"`
	Module      string                 `json:"module" structs:"module" mapstructure:"module"`
	Path        string                 `json:"path" structs:"path" mapstructure:"path"`
	Options     map[string]interface{} `json:"options" structs:"options" mapstructure:"options"`
}

type plugin struct {
	manifest Manifest
	proc     processor.EventProcessor
}

func New(conf config.Processor) (processor.EventProcessor, error) {
	var manifest Manifest
	if err := config.BindOptions(conf.Options, &manifest); err != nil {
		return nil, err
	}

	i := interp.New(interp.Options{})
	i.Use(stdlib.Symbols)
	i.Use(imports.Symbols)

	code, err := ioutil.ReadFile(filepath.Join(manifest.Path, manifest.Module))
	if err != nil {
		return nil, err
	}

	_, err = i.Eval(string(code))
	if err != nil {
		return nil, errors.Wrap(err, "eval plugin code")
	}

	pconf, err := i.Eval(manifest.Import + ".NewConfig()")
	if err != nil {
		return nil, errors.Wrap(err, "get plugin config value")
	}

	pconfI := pconf.Interface()

	if err := config.BindOptions(manifest.Options, &pconfI); err != nil {
		return nil, errors.Wrap(err, "bind plugin configuration")
	}

	spew.Dump(pconfI)

	newFn, err := i.Eval(manifest.Import + ".New")
	if err != nil {
		return nil, errors.Wrap(err, "extract plugin New()")
	}

	newArgs := []reflect.Value{reflect.ValueOf(pconfI)}

	vProc := newFn.Call(newArgs)

	proc, ok := vProc[0].Interface().(processor.EventProcessor)
	if !ok {
		return nil, errors.New("plugin is not a processor")
	}

	spew.Dump(proc)

	return &plugin{
		manifest: manifest,
		proc:     proc,
	}, nil
}

func (p *plugin) Process(evt *event.Event) ([]*event.Event, error) {
	return p.proc.Process(evt)
}
