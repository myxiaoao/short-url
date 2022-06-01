package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"short-url/model"
	"short-url/utils"
	"strconv"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

func redirect(c *fiber.Ctx) error {
	redirect := c.Params("redirect")
	url, err := model.FindByShortUrl(redirect)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not find url in DB " + err.Error(),
		})
	}
	url.Clicked += 1
	err = model.UpdateUrl(url)
	if err != nil {
		fmt.Printf("error updating: %v\n", err)
	}

	return c.Redirect(url.Redirect, fiber.StatusTemporaryRedirect)
}

func getAllUrls(c *fiber.Ctx) error {
	urls, err := model.GetAllUrls()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error getting all url links " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(urls)
}

func getUrl(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not parse id " + err.Error(),
		})
	}

	url, err := model.GetUrl(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not retrieve short from db " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(url)
}

func createUrl(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var url model.Url
	err := c.BodyParser(&url)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error parsing JSON " + err.Error(),
		})
	}

	if url.Random {
		url.Short = utils.RandomURL(8)
	}

	err = model.CreateUrl(url)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not create url in db " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(url)

}

func updateUrl(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var url model.Url

	err := c.BodyParser(&url)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not parse json " + err.Error(),
		})
	}

	err = model.UpdateUrl(url)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not update url link in DB " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(url)
}

func deleteUrl(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not parse id from url " + err.Error(),
		})
	}

	err = model.DeleteUrl(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not delete from db " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "url deleted.",
	})
}

func SetupAndListen() {
	router := fiber.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router.Get("/r/:redirect", redirect)

	router.Get("/url", getAllUrls)
	router.Get("/url/:id", getUrl)
	router.Post("/url", createUrl)
	router.Patch("/url", updateUrl)
	router.Delete("/url/:id", deleteUrl)

	err := router.Listen(":3000")
	if err != nil {
		return
	}
}
