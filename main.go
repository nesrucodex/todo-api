package main

import (
	"log"
	"todo-api/routers"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()
	api := app.Group("/api")
	routers.TodoRouter(api)
	routers.UserRouter(api)

	log.Fatal(app.Listen(":4040"))
}
