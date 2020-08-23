package http

import (
	"github.com/blushft/strana/platform/config"
	"github.com/blushft/strana/platform/logger"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
)

type Server struct {
	conf config.Server
	app  *fiber.App
}

func NewServer(conf config.Server, debug ...bool) *Server {
	isDebug := len(debug) > 0 && debug[0]

	app := fiber.New()

	app.Use(middleware.Logger())

	if !isDebug {
		app.Use(middleware.Recover())
	} else {
		logger.Log().Info("debugging enabled")
	}

	app.Get("/healthz", func(c *fiber.Ctx) {
		c.Status(200).Send("OK")
	})

	return &Server{
		conf: conf,
		app:  app,
	}
}

func (s *Server) Mount(fn func(fiber.Router) error) error {
	return fn(s.app)
}

func (s *Server) Router() fiber.Router {
	return s.app
}

func (s *Server) Start() error {
	return s.app.Listen(s.conf.HostPort())
}

func (s *Server) Shutdown() error {
	return s.app.Shutdown()
}
