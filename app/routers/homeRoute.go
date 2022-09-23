package routers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/koybigino/learn-fiber/app/controllers"
)

func HandleHomeRoute(app *fiber.App) {
	app.Get("/", controllers.Home)
}
