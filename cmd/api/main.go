package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(os.Getenv("MYSQL_CONN_STR"))
	})

	app.Listen(":8000")
}
