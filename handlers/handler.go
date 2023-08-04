package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kacioquin/golang_api/dtos"
)

func Handler() {
	fmt.Println("handlers")
}

func ShowUsersHandler(c *fiber.Ctx) error {
	user := &dtos.UserDTO{
		Name:       "Cassio",
		Profession: "Programador",
		Age:        33,
	}

	return c.JSON(user)
}

func ShowABC(c *fiber.Ctx) error {
	return c.SendString("ABC")
}
