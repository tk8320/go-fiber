package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber/handlers"
	"go-fiber/models"
	"log"
	"strconv"
	"time"
)

var CTX handlers.Context

func AddRoutes(r *fiber.App) {

	r.Get("/api/blog-post", GetAllPost)
	r.Get("/api/blog-post/:id<int>", GetPost)
	r.Delete("/api/blog-post/:id", DeletePost)
	r.Post("/api/blog-post", PostBlog)
	r.Patch("/api/blog-post/:id", UpdatePost)
	r.Post("/db/truncate", Truncate)
	CTX = handlers.InitContext()
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

	if len(data) == 0 {
		return c.Status(fiber.StatusNotFound).SendString("No Data found")
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

func Truncate(c *fiber.Ctx) error {
	var cred map[string]string
	secret_key := "Password@123"
	if err := c.BodyParser(&cred); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Unable to Truncate" + err.Error())
	}
	key, ok := cred["key"]
	if !ok {
		return c.Status(fiber.StatusBadRequest).SendString("Key is not provided")
	}

	if key != secret_key {
		return c.Status(fiber.StatusUnauthorized).SendString("Not Authorized for this transaction")
	}

	err := CTX.TruncateBlog()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Unable to Truncate" + err.Error())
	}

	return c.Status(fiber.StatusOK).SendString("Truncated Table Successfully")
}
