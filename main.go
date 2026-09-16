package main

import (
	"log"
	"os"

	"example.com/fiber-pongo2-starter/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
)

func main() {
	engine := django.New("./templates", ".html")
	engine.Reload(true) // Reload templates on every request during local development.

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/static", "./static")

	app.Get("/", handlers.Home)
	app.Get("/healthz", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	app.Get("/about", handlers.About)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8989"
	}

	log.Fatal(app.Listen("0.0.0.0:" + port))
}
