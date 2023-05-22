package main

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber/routes"
	"log"
)

func main() {

	app := fiber.New()
	routes.AddRoutes(app)

	log.Fatal(app.Listen(":8000"))

}
