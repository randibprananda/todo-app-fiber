package app

import (
	"log"
	"os"
	"todo-app-fiber/config"
	"todo-app-fiber/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func BootApp() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if portEnv := os.Getenv("PORT"); portEnv != "" {
		config.PORT = portEnv
	}

	config.BootDatabase()
	config.ConnectDatabase()
	config.RunMigration()

	app := fiber.New()

	// Cors Configs
	app.Use(cors.New(cors.Config{
		Next:             nil,
		AllowOrigins:     config.AllowOrigins,
		AllowMethods:     config.AllowMethods,
		AllowHeaders:     config.AllowHeaders,
		AllowCredentials: config.AllowCredentials,
		ExposeHeaders:    config.ExposeHeaders,
		MaxAge:           config.MaxAge,
	}))

	// Init Route
	routes.InitRoute(app)

	app.Listen(config.PORT)
}
