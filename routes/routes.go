package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	_ "go-fiber/docs"
	"go-fiber/handlers"
	"go-fiber/models"
	"log"
	"os"
	"strconv"
	"time"
)

var CTX handlers.Context

func AddRoutes(r *fiber.App) {

	r.Use(recover.New())
	r.Use(cors.New())
	r.Get("/swagger/*", swagger.HandlerDefault)
	r.Get("/api/blog-post", GetAllPost)
	r.Get("/api/blog-post/:id<int>", GetPost)
	r.Delete("/api/blog-post/:id", DeletePost)
	r.Post("/api/blog-post", PostBlog)
	r.Patch("/api/blog-post/:id", UpdatePost)
	r.Post("/db/truncate", Truncate)
	CTX = handlers.InitContext()
}

// PostBlog godoc
// @Create a Blog Post.
// @Description Create A Blog Post.
// @Tags root
// @Accept json
// @Param data body models.BlogPatch true "Request Body"
// @Produce json
// @Success 200 {object} string "Created Successfully"
// @failure 500 {object} string "Internal Server Error"
// @failure 500 {object} string "Internal Server Error"
// @Router /api/blog-post [post]
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

// GetAllPost godoc
// @Summary	List all Blog Post.
// @Description List all blog post.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {array} models.Blog
// @failure 500 {object} string "Internal Server Error"
// @failure 404 {object} string "Not Found"
// @Router /api/blog-post [get]
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

// GetPost godoc
// @Summary	Get Post by Id.
// @Description Get Post by ID.
// @Tags root
// @Accept */*
// @Produce json
// @Param id  path int true "Post ID"
// @Success 200 {object} models.Blog
// @failure 500 {object} string "Internal Server Error"
// @failure 404 {object} string "Not Found"
// @Router /api/blog-post/{id} [get]
func GetPost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	data, err := CTX.GetBlogByID(int64(id))
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.Status(fiber.StatusOK).JSON(data)
}

// DeletePost godoc
// @Summary	Delete Post by Id.
// @Description Delete Post by ID.
// @Tags root
// @Accept */*
// @Produce json
// @Param id  path int true "Post ID"
// @Success 200 {object} string "Deleted Successfully"
// @failure 500 {object} string "Internal Server Error"
// @failure 404 {object} string "Not Found"
// @Router /api/blog-post/{id} [delete]
func DeletePost(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))
	err := CTX.DeleteBlogByID(int64(id))

	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.Status(fiber.StatusOK).SendString("Deleted Successfully")
}

// UpdatePost godoc
// @Summary	Update Post by Id.
// @Description Update Post by ID.
// @Tags root
// @Accept */*
// @Produce json
// @Param id  path int true "Post ID"
// @Param data body models.BlogPatch true "Request Body"
// @Success 200 {object} string "Updated Successfully"
// @failure 500 {object} string "Internal Server Error"
// @failure 404 {object} string "Not Found"
// @Router /api/blog-post/{id} [patch]
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
	secret_key := os.Getenv("SECRET_KEY")
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
