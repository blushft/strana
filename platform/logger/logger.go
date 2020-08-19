package logger

import (
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/apex/log/handlers/json"
	"github.com/apex/log/handlers/multi"
	"github.com/apex/log/handlers/text"
	"github.com/blushft/strana/platform/config"
)

var _global *Logger

func init() {
	Init(config.DefaultLoggerConfig())
	_global = New()
}

func Log() *Logger {
	return _global
}

type Fields map[string]interface{}

func (f Fields) Fields() log.Fields {
	return log.Fields(f)
}

type Logger struct {
	log.Interface
}

func Init(conf config.Logger) {
	configure(conf)
}

func New() *Logger {
	return &Logger{
		Interface: log.Log,
	}
}

func Copy(l log.Interface) *Logger {
	return &Logger{
		Interface: l,
	}
}

func (l *Logger) WithFields(f log.Fielder) *Logger {
	return Copy(l.Interface.WithFields(f))
}

func (l *Logger) Mount(fn func(*Logger)) {
	fn(l)
}

func configure(c config.Logger) {
	hdlrs := handlers(c)
	if len(hdlrs) > 1 {
		log.SetHandler(multi.New(hdlrs...))

		return
	}

	log.SetHandler(hdlrs[0])
}

func handlers(c config.Logger) []log.Handler {
	hdlrs := make([]log.Handler, 0, len(c.Outputs))
	for _, o := range c.Outputs {
		hdlrs = append(hdlrs, configHandler(o))
	}

	return hdlrs
}

func configHandler(c config.Output) log.Handler {
	switch c.Type {
	case "text":
		return text.New(os.Stderr)
	case "cli":
		return cli.New(os.Stderr)
	case "json":
		return json.New(os.Stderr)
	default:
		return text.New(os.Stderr)
	}
}
