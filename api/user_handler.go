package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sendistephen/room-booking/types"
)

func HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "Stephen",
		LastName:  "Ssendikadiwa",
	}
	return c.JSON(u)
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON("Single user")
}
