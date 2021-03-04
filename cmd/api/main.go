package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gopher-lib/config"
	"github.com/mlukasik-dev/test/internal/appconfig"

	_ "github.com/lib/pq"
)

func main() {
	var conf appconfig.AppConfig
	err := config.Load(&conf, "configs/config.yaml")

	s := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", os.Getenv("RDS_USERNAME"), os.Getenv("RDS_PASSWORD"), os.Getenv("RDS_HOSTNAME"), os.Getenv("RDS_PORT"), os.Getenv("RDS_DB_NAME"))
	db, err := sql.Open("postgres", s)
	if err != nil {
		log.Printf("failed to connect to postgres db: %v", err)
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

	addr := fmt.Sprintf(":%d", conf.Port)
	app.Listen(addr)
}
