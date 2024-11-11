package handlers

import "github.com/gofiber/fiber/v2"
import "todo-api/models"

var users []models.User = []models.User{
	{ID: 1, Name: "John Doe", Email: "john@example.com", Password: "password1"},
	{ID: 2, Name: "Jane Doe", Email: "jane@example.com", Password: "password2"},
	{ID: 3, Name: "Alice Doe", Email: "alice@example.com", Password: "password3"},
}

func GetUsers(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(models.ServerResponse{Success: true, Message: "Users retrieved successfully", Data: users})
}
