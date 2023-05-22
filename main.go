package main

import (
	"github.com/gofiber/fiber/v2"
	_ "go-fiber/docs"
	"go-fiber/routes"
	"log"
)

// @title Fiber Swagger Example API
// @version 2.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host go-fiber-app.onrender.com
// @BasePath /
// @schemes http
func main() {

	app := fiber.New()
	routes.AddRoutes(app)

	log.Fatal(app.Listen(":8000"))

}
