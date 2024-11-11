package routers

import (
	"github.com/gofiber/fiber/v2"
	"todo-api/handlers"
)

func TodoRouter(api fiber.Router) {
	todo := api.Group("/todos")
	todo.Get("/", handlers.GetTodos)
	todo.Get("/:id", handlers.GetTodoById)
	todo.Post("/", handlers.CreateTodo)
}
