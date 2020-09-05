package controller

import (
	"net/http"

	"github.com/gofiber/fiber"
)

func (a *App) routes(rtr fiber.Router) error {
	api := rtr.Group("/api")
	cg := api.Group("/controller")

	cg.Get("/config", func(ctx *fiber.Ctx) {
		if err := ctx.JSON(a.conf); err != nil {
			ctx.Status(http.StatusInternalServerError).Send(err)
		}
	})

	return nil
}
