package main

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/gofiber/template/html"
)

func main() {
	// Create a new engine
	engine := html.New("./views", ".html")

	// Or from an embedded system
	// See github.com/gofiber/embed for examples
	// engine := html.NewFileSystem(http.Dir("./views", ".html"))

	// Pass the engine to the Views
	app := fiber.New(&fiber.Settings{
		Views: engine,
	})

	// compression with default config
	app.Use(middleware.Compress())

	app.Get("/", func(c *fiber.Ctx) {
		// Render index
		c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	})

	app.Get("/layout", func(c *fiber.Ctx) {
		// Render index within layouts/main
		c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		}, "layouts/main")
	})

	app.Listen(3000)
}
