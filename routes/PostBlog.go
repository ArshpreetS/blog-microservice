package routes

import (
	"context"
	"fmt"

	"github.com/arshpreets/blog_service/database"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type blog struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func getNextID(b *mongo.Collection) int {
	findOptions := options.FindOne()
	findOptions.SetSort(bson.D{{"id", -1}})

	lastBlogEntry := &blog{}

	if err := b.FindOne(context.TODO(), bson.M{}, findOptions).Decode(&lastBlogEntry); err != nil {
		if err == mongo.ErrNoDocuments {
			lastBlogEntry = &blog{ID: 0, Title: "hey", Content: "hey"}
		} else {
			panic(err)
		}
	}
	return lastBlogEntry.ID + 1
}

func PostBlog(c *fiber.Ctx) error {

	db_client := database.Get_dbClient()
	defer db_client.Disconnect(context.Background())
	blogCollections := db_client.Database("myBlogs").Collection("blogList")

	nextID := getNextID(blogCollections)

	nextBlog := &blog{}

	if err := c.BodyParser(&nextBlog); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Some issues with unmarshalling json data"})
	}
	nextBlog.ID = nextID
	_, err := blogCollections.InsertOne(context.TODO(), nextBlog)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Issue with sending data to database"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": "Cool"})
}
