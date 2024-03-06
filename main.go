package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
)

func main() {
	listenAddress := flag.String("listenAddress", ":5000", "The listening address of the server api")
	flag.Parse()
	app := fiber.New()
	apiv1 := app.Group("/api/v1")
	apiv1.Get("/foo", handleFoo)

	app.Listen(*listenAddress)
}

func handleFoo(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "hello foo"})
}
