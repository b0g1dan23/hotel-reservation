package api

import (
	"github.com/b0g1dan23/hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
)

func HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "Bogdan",
		LastName:  "Stevanovic",
	}
	return c.JSON(u)
}

func HandleGetUserByID(c *fiber.Ctx) error {
	return c.JSON("James")
}
