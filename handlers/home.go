package handlers

import "github.com/gofiber/fiber/v2"

// Home renders the application's home page.
func Home(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Fiber + Pongo2",
		"Items": []string{"Fiber", "Pongo2", "Server-side rendering"},
	})
}
