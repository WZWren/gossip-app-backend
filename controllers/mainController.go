package controllers

import (
	"fmt"

	"github.com/WZWren/gossip-app-backend/database"
	"github.com/WZWren/gossip-app-backend/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Healthcheck(c *fiber.Ctx) error {
	return c.SendString("Healthcheck OK!")
}

/****
 * this is the basic body of every controller func - fetching the
 * data sent in by the request
 *
 * var data map[string]string
 *
 *	if err := c.BodyParser(&data); err != nil {
 *		return err
 *	}
 *  ...
 * **/

func Register(c *fiber.Ctx) error {
	//this creates a "array" that maps a string key to a string value
	//think this like a json
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var checkuser models.User
	database.DB.Where("name = ?", data["user_name"]).First(&checkuser)

	if checkuser.Id != 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Username already in use.",
		})
	}

	// []byte takes a string and maps it to a byte array
	password, _ := bcrypt.GenerateFromPassword([]byte(data["user_pass"]), bcrypt.DefaultCost)

	user := models.User{
		Name:     data["user_name"],
		Password: password,
	}

	database.DB.Create(&user)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("name = ?", data["user_name"]).First(&user)

	// user.Id is init to 0
	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User not found.",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["user_pass"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		fmt.Printf("err: %v\n", err)
		return c.JSON(fiber.Map{
			"message": "Incorrect Password.",
		})
	}

	// this only reaches if all prev error checks are passed
	return c.JSON(user)
}
