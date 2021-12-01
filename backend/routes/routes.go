package routes

import (
	"github.com/gahtoe/pfol/database"
	"github.com/gahtoe/pfol/models"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	db := database.DB.Db
	user := new(models.User)
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	db.Create(&user)

	return c.Status(200).JSON(user)
}

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}
	database.DB.Db.Find(&users)

	return c.Status(200).JSON(users)
}
