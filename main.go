package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"

	"github.com/koybigino/learn-fiber/app/routers"
)

func main() {

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	routers.HandleHomeRoute(app)
	routers.HandlePostRoute(app)

	log.Fatal(app.Listen(":3000"))
}
