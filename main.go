package main

import (
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sendistephen/room-booking/api"
)

func main() {
	listenAddr := flag.String("listenAddr", ":4000", "The listening address of the server api")
	flag.Parse()

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
