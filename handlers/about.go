package handlers

import "github.com/gofiber/fiber/v2"

// Home renders the application's home page.
func About(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "About Us",
		"Items": []string{"Fiber", "Pongo2", "Server-side rendering"},
	})
}
