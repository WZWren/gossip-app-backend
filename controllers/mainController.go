package controllers

import "github.com/gofiber/fiber/v2"

func Healthcheck(c *fiber.Ctx) error {
	return c.SendString("Healthcheck OK!")
}
