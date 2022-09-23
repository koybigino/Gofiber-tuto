package controllers

import (
	"fmt"
	"strconv"

	"github.com/koybigino/learn-fiber/app/database"
	"github.com/koybigino/learn-fiber/app/models"

	"github.com/gofiber/fiber/v2"
)

var db = database.ConnectionDB()

func GetAllPosts(c *fiber.Ctx) error {
	var posts []models.Post
	var count int64
	db.Find(&posts).Count(&count)
	return c.JSON(fiber.Map{
		"posts":         posts,
		"numberOfPosts": count,
	})
}

func GetPostByID(c *fiber.Ctx) error {
	strId := c.Params("id")

	id, err := strconv.Atoi(strId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error to parse the id parameter",
		})
	}

	post := new(models.Post)

	if err := db.First(post, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   err.Error(),
			"message": fmt.Sprintf("Any Post correspond to this id = %d", id),
		})
	}
	return c.JSON(fiber.Map{
		"post": post,
	})
}

func CreatePost(c *fiber.Ctx) error {
	newPost := new(models.Post)

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

}

func UpdatePost(c *fiber.Ctx) error {
	strid := c.Params("id")
	updatePost := new(models.Post)

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

	post := new(models.Post)

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

}

func DeletePost(c *fiber.Ctx) error {
	strid := c.Params("id")

	id, err := strconv.Atoi(strid)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error to parse the id parameter",
		})
	}

	post := new(models.Post)

	if err := db.First(post, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   err.Error(),
			"message": fmt.Sprintf("Any Post correspond to this id = %d", id),
		})
	}

	db.Delete(post)

	return c.SendStatus(fiber.StatusNoContent)
}
