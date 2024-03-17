package routes

import (
	"github.com/discrete-taburetka/vpn.telegrambot/models"
	"github.com/discrete-taburetka/vpn.telegrambot/database"
	"github.com/gofiber/fiber/v2"
)

// Hello
func Hello(c *fiber.Ctx) error {
	return c.SendString("fiber")
}

// AddUser
func AddUser(c *fiber.Ctx) error {
	new_user := new(models.Tg_users)
	if err := c.BodyParser(new_user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Db.Create(&new_user)

	return c.Status(200).JSON(new_user)
}

// AllUsers
func AllUser(c *fiber.Ctx) error {
	users := []models.Tg_users{}
	database.DB.Db.Find(&users)

	return c.Status(200).JSON(users)
}

// Delete
func Delete(c *fiber.Ctx) error {
	users := []models.Tg_users{}
	tg_user_name := new(models.Tg_users)
	if err := c.BodyParser(tg_user_name); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.DB.Db.Where("tg_user_name = ?", tg_user_name.Tg_user_name).Delete(&users)

	return c.Status(200).JSON("deleted")
}