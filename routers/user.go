package routers

import (
	"github.com/gofiber/fiber/v2"
	"todo-api/handlers"
)

func UserRouter(api fiber.Router) {
	userGroup := api.Group("/users")
	userGroup.Get("/", handlers.GetUsers)
}
