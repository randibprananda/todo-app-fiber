package controllers

import (
	"log"
	"todo-app-fiber/database"
	"todo-app-fiber/models"
	"todo-app-fiber/request"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func CreateTodo(c *fiber.Ctx) error {
	todoReq := request.TodoCreateRequest{}

	// Parse request body
	if errParse := c.BodyParser(&todoReq); errParse != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Fail to parsing data",
			"error":   errParse.Error(),
		})
	}

	// Validation request data
	validate := validator.New()

	if errValidate := validate.Struct(&todoReq); errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Some data is not valid",
			"error":   errValidate.Error(),
		})
	}

	todo := models.Todo{}
	todo.Name = todoReq.Name
	todo.IsComplete = todoReq.IsComplete

	if todoReq.Note != "" {
		todo.Note = todoReq.Note
	}

	if errDb := database.DB.Create(&todo).Error; errDb != nil {
		log.Println("todo.controller.go => CreateTodo :: ", errDb)

		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Todo created successfully",
		"data":    todo,
	})

}

func GetAllTodo(c *fiber.Ctx) error {
	todos := []models.Todo{}

	err := database.DB.Find(&todos).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data transmited",
		"data":    todos,
	})
}

func GetTodoByID(c *fiber.Ctx) error {
	return nil
}

func UpdateTodoByID(c *fiber.Ctx) error {
	return nil
}

func DeleteTodoById(c *fiber.Ctx) error {
	return nil
}
