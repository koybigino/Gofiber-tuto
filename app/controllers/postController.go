package controllers

import (
	"fmt"
	"strconv"

	"github.com/koybigino/learn-fiber/app/database"
	"github.com/koybigino/learn-fiber/app/models"
	"github.com/koybigino/learn-fiber/app/validation"

	"github.com/gofiber/fiber/v2"
)

var db = database.ConnectionDB()

func GetAllPosts(c *fiber.Ctx) error {
	var posts []models.Post
	var postsResponse []models.PostResponse
	var count int64
	db.Find(&posts).Count(&count)

	for _, post := range posts {
		pr := new(models.PostResponse)
		models.ParseToResponse(post, pr)
		postsResponse = append(postsResponse, *pr)
	}
	return c.JSON(fiber.Map{
		"posts":         postsResponse,
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
	postResponse := new(models.PostResponse)

	if err := db.First(post, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   err.Error(),
			"message": fmt.Sprintf("Any Post correspond to this id = %d", id),
		})
	}

	models.ParseToResponse(*post, postResponse)

	return c.JSON(fiber.Map{
		"post": postResponse,
	})
}

func CreatePost(c *fiber.Ctx) error {
	newPost := new(models.Post)
	newPostRequest := new(models.PostRequest)
	newPostResponse := new(models.PostResponse)

	if err := c.BodyParser(newPostRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	errors := validation.ValidateStruct(newPostRequest)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": errors,
		})
	}

	models.ParseToPost(newPost, *newPostRequest)

	if err := db.Create(newPost).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Error to createthe element !",
		})
	}

	models.ParseToResponse(*newPost, newPostResponse)

	return c.JSON(fiber.Map{
		"post": newPostResponse,
	})

}

func UpdatePost(c *fiber.Ctx) error {
	strid := c.Params("id")
	updatePost := new(models.Post)
	updatePostRequest := new(models.PostRequest)
	updatePostResponse := new(models.PostResponse)

	id, err := strconv.Atoi(strid)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error to parse the id parameter",
		})
	}

	if err := c.BodyParser(updatePostRequest); err != nil {
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

	models.ParseToPost(updatePost, *updatePostRequest)

	db.Save(updatePost)

	models.ParseToResponse(*updatePost, updatePostResponse)

	return c.JSON(fiber.Map{
		"post": updatePostResponse,
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
