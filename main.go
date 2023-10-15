package main

import (
	"todo-app-fiber/config"
	"todo-app-fiber/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.BootApp()
	app := fiber.New()

	// Init Route
	routes.InitRoute(app)

	app.Listen(":8000")
}
