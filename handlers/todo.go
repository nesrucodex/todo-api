package handlers

import (
	"strconv"
	"todo-api/models"

	"github.com/gofiber/fiber/v2"
)

var todos = []models.Todo{}

func findTodoById(id int) *models.Todo {
	for _, t := range todos {
		if t.ID == id {
			return &t
		}
	}
	return nil
}

func GetTodos(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(models.ServerResponse{
		Success: true,
		Message: "Todos retrieved successfully",
		Data:    todos,
	})
}

func GetTodoById(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.ServerResponse{
			Success: false,
			Message: "Invalid ID",
			Data:    nil,
		})
	}

	var todo *models.Todo = findTodoById(id)

	if todo == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(models.ServerResponse{
			Success: false,
			Message: "Todo not found",
			Data:    nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(models.ServerResponse{
		Success: true,
		Message: "Todo retrieved successfully",
		Data:    todo,
	})

}

func CreateTodo(ctx *fiber.Ctx) error {

	var todo = &models.Todo{}

	if err := ctx.BodyParser(todo); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.ServerResponse{
			Success: false,
			Message: "Invalid request body",
			Data:    nil,
		})
	}

	todo.ID = len(todos) + 1

	if todo.Body == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.ServerResponse{
			Success: false,
			Message: "Todo body is required",
			Data:    nil,
		})
	}

	todos = append(todos, *todo)
	return ctx.Status(fiber.StatusCreated).JSON(models.ServerResponse{
		Success: true,
		Message: "Todo created successfully",
		Data:    todo,
	})

}

func UpdateTodo(ctx *fiber.Ctx) error {
	type UpdateTodoDTO struct {
		Body        *string `json:"body"`
		IsCompleted *bool   `json:"isCompleted"`
	}
	var updateTodoDto UpdateTodoDTO
	if err := ctx.BodyParser(&updateTodoDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.ServerResponse{
			Success: false,
			Message: "Invalid request body",
			Data:    nil,
		})

	}

	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.ServerResponse{
			Success: false,
			Message: "Invalid ID",
			Data:    nil,
		})
	}

	for i, todo := range todos {
		if todo.ID == id {
			if updateTodoDto.Body != nil && *updateTodoDto.Body != "" {
				todos[i].Body = *updateTodoDto.Body
			}
			if updateTodoDto.IsCompleted != nil {
				todos[i].IsCompleted = *updateTodoDto.IsCompleted
			}
			return ctx.Status(fiber.StatusOK).JSON(models.ServerResponse{
				Success: true,
				Message: "Todo updated successfully",
				Data:    todos[i],
			})
		}
	}

	return ctx.Status(fiber.StatusNotFound).JSON(models.ServerResponse{
		Success: false,
		Message: "Todo not found",
		Data:    nil,
	})
}

func DeleteTodo(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.ServerResponse{
			Success: false,
			Message: "Invalid ID",
			Data:    nil,
		})
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)

			return ctx.Status(fiber.StatusOK).JSON(models.ServerResponse{
				Success: true,
				Message: "Todo deleted successfully",
				Data:    todos[i],
			})
		}
	}

	return ctx.Status(fiber.StatusNotFound).JSON(models.ServerResponse{
		Success: false,
		Message: "Todo not found",
		Data:    nil,
	})
}
