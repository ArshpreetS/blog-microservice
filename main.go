package main

import (
	"log"

	"github.com/arshpreets/blog_service/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func setRoutes(app *fiber.App) {

	app.Get("/allblogs", routes.GetAllBlogs)
	app.Post("/postblog", routes.PostBlog)
	app.Get("/getblog/:id", routes.GetBlog)
}

func main() {
	app := fiber.New()

	// Create the routes
	setRoutes(app)
	app.Use(logger.New())
	log.Fatal(app.Listen(":3000"))
}
