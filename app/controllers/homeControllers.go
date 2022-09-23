package controllers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/koybigino/learn-fiber/app/models"
)

func Home(c *fiber.Ctx) error {
	var posts []models.Post
	var count int64

	db.Find(&posts).Count(&count)

	return c.Render("index", fiber.Map{
		"hello": "Hello World !",
		"posts": posts,
		"count": count,
	})
}
