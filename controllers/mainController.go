package controllers

import "github.com/gofiber/fiber/v2"

func Healthcheck(c *fiber.Ctx) error {
	return c.SendString("Healthcheck OK!")
}

func Register(c *fiber.Ctx) error {
	//this creates a "array" that maps a string key to a string value
	//think this like a json
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	return c.JSON(data)
}
