package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gopher-lib/config"
	"github.com/mlukasik-dev/test/internal/appconfig"

	_ "github.com/lib/pq"
)

func main() {
	var conf appconfig.AppConfig
	err := config.Load(&conf, "configs/config.yaml")

	db, err := sql.Open("postgres", conf.PostgresConnStr)
	if err != nil {
		log.Fatalf("failed to connect to postgres db: %v", err)
	}
	defer db.Close()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		var message string
		err := db.QueryRowContext(c.Context(), "SELECT message FROM messages LIMIT 1").Scan(&message)
		if err != nil {
			return c.Status(http.StatusInternalServerError).SendString(err.Error())
		}
		return c.SendString(message)
	})

	app.Listen(":8080")
}
