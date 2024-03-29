package main

import (

		"fmt"
		"github.com/gofiber/fiber/v2"
)


func main() {
	fmt.Println("Hello there buddy!");
	app := fiber.New()

	// give response when at /
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success":  true,
			"message": "You are at the endpoint 😉",
		})
	})

	// Listen on server 8000 and catch error if any
	err := app.Listen(":8000")

	// handle error
	if err != nil {
		panic(err)
	}
}