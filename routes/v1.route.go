package routes

import (
	"todo-app-fiber/controllers"

	"github.com/gofiber/fiber/v2"
)

func v1Route(app *fiber.App) {
	v1 := app.Group("/api/v1")

	// Todo
	todo := v1.Group("/todo")
	todo.Post("/", controllers.CreateTodo)
	todo.Get("/", controllers.GetAllTodo)
	todo.Get("/:id", controllers.GetTodoByID)
	todo.Patch("/:id", controllers.UpdateTodoByID)
	todo.Delete("/:id", controllers.DeleteTodoById)
}
