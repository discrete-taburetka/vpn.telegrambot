package main

import (
	"log"

	"github.com/discrete-taburetka/vpn.telegrambot/database"
	"github.com/discrete-taburetka/vpn.telegrambot/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/hello", routes.Hello)
	app.Get("/allusers", routes.AllUser)
	app.Post("/adduser", routes.AddUser)
	app.Delete("/user", routes.Delete)
}


func main() {
	database.ConnectDb()

	app := fiber.New()

	setUpRoutes(app)

	app.Use(cors.New())

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	app.Static("/", "static/")

	log.Fatal(app.Listen(":8080"))
}
