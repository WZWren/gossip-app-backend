package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	_, err := gorm.Open(mysql.Open("root:rootwalla@/gossip"), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("I LIVE LADS!")
	})

	app.Listen(":8000")
}
