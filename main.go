package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Post struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Published bool      `json:"published"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func main() {

	dsn := "host=localhost user=postgres password=Bielem@*01 dbname=fiber_db port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Database connection Error")
	}

	fmt.Println("Connection succed !")

	if err := db.AutoMigrate(&Post{}); err != nil {
		panic("Error ro create the table")
	}

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		var posts []Post
		var count int64

		db.Find(&posts).Count(&count)

		return c.Render("index", fiber.Map{
			"hello": "Hello World !",
			"posts": posts,
			"count": count,
		})
	})

	// Get all posts
	app.Get("/posts", func(c *fiber.Ctx) error {
		var posts []Post
		var count int64
		db.Find(&posts).Count(&count)
		return c.JSON(fiber.Map{
			"posts":         posts,
			"numberOfPosts": count,
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

		post := new(Post)

		if err := db.First(post, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":   err.Error(),
				"message": fmt.Sprintf("Any Post correspond to this id = %d", id),
			})
		}
		return c.JSON(fiber.Map{
			"post": post,
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

		if err := db.Create(newPost).Error; err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   err.Error(),
				"message": "Error to createthe element !",
			})
		}

		return c.JSON(fiber.Map{
			"post": newPost,
		})

	})

	// Upadate a post
	app.Put("/posts/:id", func(c *fiber.Ctx) error {
		strid := c.Params("id")
		updatePost := new(Post)

		id, err := strconv.Atoi(strid)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "error to parse the id parameter",
			})
		}

		if err := c.BodyParser(updatePost); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		post := new(Post)

		if err := db.First(post, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":   err.Error(),
				"message": fmt.Sprintf("Any Post correspond to this id = %d", id),
			})
		}

		updatePost.Id = id

		db.Save(updatePost)

		return c.JSON(fiber.Map{
			"post": updatePost,
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

		post := new(Post)

		if err := db.First(post, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":   err.Error(),
				"message": fmt.Sprintf("Any Post correspond to this id = %d", id),
			})
		}

		db.Delete(post)

		return c.SendStatus(fiber.StatusNoContent)
	})

	// app.Static("/", "./public/index.html")

	log.Fatal(app.Listen(":3000"))
}
