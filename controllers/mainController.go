package controllers

import (
	"strconv"
	"time"

	"github.com/WZWren/gossip-app-backend/database"
	"github.com/WZWren/gossip-app-backend/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret"

// basic route of the backend, allows us to check if backend is up.
// GET request.
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

// Register route. This will check if username is already in use before
// storing the new user into the database.
// POST request.
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
	password, _ := bcrypt.GenerateFromPassword(
		[]byte(data["user_pass"]), bcrypt.DefaultCost)

	user := models.User{
		Name:     data["user_name"],
		Password: password,
	}

	database.DB.Create(&user)

	return c.JSON(user)
}

// login route. This will return a cookie on auth success.
// POST request.
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

	if err := bcrypt.CompareHashAndPassword(
		user.Password, []byte(data["user_pass"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Incorrect Password.",
		})
	}

	// creates a precursor to a cookie that expires in 24 hours
	claims := jwt.NewWithClaims(
		jwt.SigningMethodHS256, jwt.StandardClaims{
			Issuer:    strconv.Itoa(int(user.Id)),
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Could not log in.",
		})
	}

	cookie := fiber.Cookie{
		Name:     "user-login-gossip",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	// this only reaches if all prev error checks are passed
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Success.",
	})
}

// User Auth route, using the HTTP cookie. This is the actual route
// the frontend will use to auth a login - the login route will generate
// the cookie for the User func to fetch. GET request.
func User(c *fiber.Ctx) error {
	cookie := c.Cookies("user-login-gossip")

	token, err := jwt.ParseWithClaims(
		cookie, &jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "User is not logged in.",
		})
	}

	// map the token claims to StandardClaims so we can access the Issuer call
	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	database.DB.Where("id = ?", claims.Issuer).First(&user)

	return c.JSON(user)
}

// Logout route. Logout is done by expiring the cookie on the frontend.
// GET request.
func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "user-login-gossip",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Logout successful.",
	})
}

// PostThread route. Sends a Thread into the database.
// POST request.
func PostThread(c *fiber.Ctx) error {
	var thread models.Thread

	if err := c.BodyParser(&thread); err != nil {
		return err
	}

	database.DB.Create(&thread)

	return c.JSON(thread)
}

// GetThread route. Grants a JSON Object with all the threads to display.
// GET request.
func GetThreads(c *fiber.Ctx) error {
	var threads []models.Thread
	database.DB.Find(&threads)
	return c.JSON(threads)
}

// PostComment route. Sends a Comment into the database.
// POST request.
func PostComment(c *fiber.Ctx) error {
	var comment models.Comment

	if err := c.BodyParser(&comment); err != nil {
		return err
	}
	time := time.Now().Unix()

	comment.DateCreated = time
	comment.DateUpdated = time

	database.DB.Create(&comment)

	return c.JSON(comment)
}

// GetComments route. Grants a JSON Object with all the comments to display.
// This specifically only looks for the data corresponding to the thread id
// given.
// GET request.
func GetComments(c *fiber.Ctx) error {
	// data should only consist of "thread_id": integer.
	var data map[string]int

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var comments []models.Comment
	result := database.DB.Where("thread_id = ?", data["thread_id"]).Find(&comments)

	// get the thread the comment belongs to, to update one of its values.
	// this is done here, instead of postcomment as this is always called
	// immediately after postcomment.
	database.DB.Model(&models.Thread{}).Where(
		"id = ?", data["thread_id"]).Update("comment_no", uint(result.RowsAffected))

	return c.JSON(comments)
}
