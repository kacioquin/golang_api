package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kacioquin/golang_api/handlers"
)

var app = fiber.New()

func Router() {
	appV1 := app.Group("/v1")

	mountUserRoutes(appV1)

	app.Listen(":3000")

}

func mountUserRoutes(app fiber.Router) {
	userRoute := app.Group("/user")
	userRoute.Get("/", handlers.ShowUsersHandler)
	userRoute.Get("/abc/", handlers.ShowABC)
}
