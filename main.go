package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Post struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Published bool   `json:"published"`
}

func main() {
	app := fiber.New()

	var posts []Post

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello Wolrd !")
	})

	// Get all posts
	app.Get("/posts", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"posts": posts,
		})
	})

	// Get a specific post
	app.Get("/posts/:id", func(c *fiber.Ctx) error {
		strId := c.Params("id")

		id, err := strconv.Atoi(strId)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "error to parse the id parameter",
			})
		}

		for _, post := range posts {
			if post.Id == id {
				return c.JSON(fiber.Map{
					"posts": post,
				})
			}
		}

		return c.JSON(fiber.Map{
			"massage": fmt.Sprintf("Any Post correspond to this id = %d", id),
		})
	})

	// Create a Post
	app.Post("/posts", func(c *fiber.Ctx) error {
		newPost := new(Post)

		if err := c.BodyParser(newPost); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		newPost.Id = rand.Intn(100000)

		posts = append(posts, *newPost)

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"posts": newPost,
		})
	})

	// Upadate a post
	app.Put("/posts/:id", func(c *fiber.Ctx) error {
		strid := c.Params("id")

		id, err := strconv.Atoi(strid)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "error to parse the id parameter",
			})
		}

		updatePost := new(Post)

		for i, post := range posts {
			if post.Id == id {
				posts = append(posts[:i], posts[i+1:]...)

				if err := c.BodyParser(updatePost); err != nil {
					return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
						"error": err.Error(),
					})
				}

				updatePost.Id = id
				posts = append(posts, *updatePost)
				return c.JSON(fiber.Map{
					"posts": updatePost,
				})

			}
		}

		return c.JSON(fiber.Map{
			"massage": fmt.Sprintf("Any Post correspond to this id = %d", id),
		})

	})

	// Delete a post
	app.Delete("/posts/:id", func(c *fiber.Ctx) error {
		strid := c.Params("id")

		id, err := strconv.Atoi(strid)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "error to parse the id parameter",
			})
		}

		for i, post := range posts {
			if post.Id == id {
				posts = append(posts[:i], posts[i+1:]...)

				return c.SendStatus(fiber.StatusNoContent)

			}
		}

		return c.JSON(fiber.Map{
			"massage": fmt.Sprintf("Any Post correspond to this id = %d", id),
		})
	})

	// app.Static("/", "./public/index.html")

	log.Fatal(app.Listen(":3000"))
}
