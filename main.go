package main

import (
	"log"
	"main/handlers"
	"main/models"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

var CTX handlers.Context

func main() {

	app := fiber.New()

	app.Get("/api/blog-post", GetAllPost)
	app.Get("/api/blog-post/:id<int>", GetPost)
	app.Delete("/api/blog-post/:id", DeletePost)
	app.Post("/api/blog-post", PostBlog)
	app.Patch("/api/blog-post/:id", UpdatePost)
	CTX = handlers.InitContext()
	log.Fatal(app.Listen(":8000"))

}

func PostBlog(c *fiber.Ctx) error {
	var p models.Blog

	if err := c.BodyParser(&p); err != nil {
		return err
	}
	p.Created_At = time.Now()
	p.Updated_At = time.Now()

	// Database = append(Database, p)

	err := CTX.CreateBlog(p)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(err)
	}
	return c.Status(fiber.StatusCreated).JSON(map[string]string{"data": "Created Successfully"})
}

func GetAllPost(c *fiber.Ctx) error {
	data, err := CTX.GetAllBlogs()
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.Status(fiber.StatusOK).JSON(data)
}

func GetPost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	data, err := CTX.GetBlogByID(int64(id))
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.Status(fiber.StatusOK).JSON(data)
}

func DeletePost(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))
	err := CTX.DeleteBlogByID(int64(id))

	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.Status(fiber.StatusOK).SendString("Deleted Successfully")
}

func UpdatePost(c *fiber.Ctx) error {
	var p models.BlogPatch
	if err := c.BodyParser(&p); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Params("id"))

	err := CTX.UpdateBlog(p, int64(id))

	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.Status(fiber.StatusOK).SendString("Updated Successfully")
}
