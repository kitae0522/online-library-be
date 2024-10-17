package controller

import (
	"github.com/gofiber/fiber/v2"

	"github.com/kitae0522/online-library-be/internal/model"
)

func EnrollRouter(app *fiber.App, dbconn *model.PrismaClient) {
	apiRouter := app.Group("/api")
	initAuthRouter(apiRouter, initAuthDI(dbconn))
	initUserRouter(apiRouter, initUserDI(dbconn))

	apiRouter.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"message": "pong",
		})
	})
}
