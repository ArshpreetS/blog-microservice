package routes

import (
	"context"

	"github.com/arshpreets/blog_service/database"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type blogJSON struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func GetAllBlogs(c *fiber.Ctx) error {

	db_client := database.Get_dbClient()
	defer db_client.Disconnect(context.TODO())

	if err := db_client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	blog_collections := db_client.Database("myBlogs").Collection("blogList")

	// searching for data
	cursor, err := blog_collections.Find(context.Background(), bson.M{}, options.Find().SetProjection(bson.M{"_id": 0}))
	if err != nil {
		panic(err)
	}

	defer cursor.Close(context.Background())

	allBlogs := []*blogJSON{}
	for cursor.Next(context.Background()) {
		currBlog := &blogJSON{}
		if err := cursor.Decode(&currBlog); err != nil {
			panic(err)
		}
		allBlogs = append(allBlogs, currBlog)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"all_blogs": allBlogs})
}
