package plugin

import (
	"path/filepath"
	"testing"

	"github.com/blushft/strana/platform/config"
	"github.com/fatih/structs"
	"github.com/stretchr/testify/suite"
)

type PluginSuite struct {
	suite.Suite
}

func TestRunPluginSuite(t *testing.T) {
	suite.Run(t, new(PluginSuite))
}

func (s *PluginSuite) TestNewPlugin() {
	fp, err := filepath.Abs("./example")
	if err != nil {
		s.FailNow("get abs path", err)
	}

	conf := config.Processor{
		Name: "myplugin",
		Type: "plugin",
		Options: structs.Map(Manifest{
			Name:        "test",
			Description: "testing",
			Version:     "0.0.1",
			Type:        "processor",
			Import:      "example",
			Module:      "example.go",
			Path:        fp,
			Options: map[string]interface{}{
				"name": "Test Plugin",
			},
		}),
	}

	plug, err := New(conf)
	if err != nil {
		s.FailNow("create plugin", err)
	}

	_, err = plug.Process(nil)
	s.Require().NoError(err)

}
