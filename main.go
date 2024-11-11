package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"todo-api/config"
	"todo-api/routers"
)

func main() {

	err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	api := app.Group("/api")
	routers.TodoRouter(api)
	routers.UserRouter(api)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", config.Config.PORT)))
}
