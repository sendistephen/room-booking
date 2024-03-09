package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sendistephen/room-booking/api"
	"github.com/sendistephen/room-booking/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb+srv://admin:admin@voicewhisperstestdb.8zqpuhq.mongodb.net/?retryWrites=true&w=majority&appName=voicewhisperstestdb"
const dbName = "room-booking"
const userCollection = "users"

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}
	listenAddr := flag.String("listenAddr", ":4000", "The listening address of the server api")
	flag.Parse()

	ctx := context.Background()
	collections := client.Database(dbName).Collection(userCollection)

	user := types.User{
		FirstName: "Stephen",
		LastName:  "Ssendikadiwa",
	}
	res, err := collections.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
	fmt.Println(client)
	app := fiber.New()
	apiv1 := app.Group("/api/v1")
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("Hello")
	})
	apiv1.Get("/user", api.HandleGetUsers)
	apiv1.Get("/:id", api.HandleGetUser)

	if err := app.Listen(*listenAddr); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
