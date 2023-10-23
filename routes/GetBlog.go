package routes

import (
	"context"
	"fmt"
	"strconv"

	"github.com/arshpreets/blog_service/database"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetBlog(c *fiber.Ctx) error {

	// creating a blog post
	id_to_get, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID provided is wrong"})
	}
	db_client := database.Get_dbClient()
	defer db_client.Disconnect(context.Background())

	blogCollections := db_client.Database("myBlogs").Collection("blogList")

	curr_blog := &blog{}
	if err := blogCollections.FindOne(context.TODO(), bson.M{"id": id_to_get}).Decode(&curr_blog); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No Such blog exists"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"blog": curr_blog})
}
