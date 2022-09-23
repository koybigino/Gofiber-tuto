package routers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/koybigino/learn-fiber/app/controllers"
)

func HandlePostRoute(app *fiber.App) {
	// Get all posts
	app.Get("/posts", controllers.GetAllPosts)

	// Get a specific post
	app.Get("/posts/:id", controllers.GetPostByID)

	// Create a Post
	app.Post("/posts", controllers.CreatePost)

	// Upadate a post
	app.Put("/posts/:id", controllers.UpdatePost)

	// Delete a post
	app.Delete("/posts/:id", controllers.DeletePost)
}
