package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
)

func main() {
	// Create a new engine
	engine := html.New("./views", ".html")
	engine.Reload(true)
	// Or from an embedded system
	// See github.com/gofiber/embed for examples
	// engine := html.NewFileSystem(http.Dir("./views", ".html"))

	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(favicon.New())
	// compression with default config
	app.Use(compress.New())

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	})

	app.Get("/layout", func(c *fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		}, "layouts/main")
	})

	app.Get("/index2", func(c *fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render("index2", fiber.Map{
			"Title": "Hello, World!",
		}, "layouts/main")
	})

	log.Fatal(app.Listen(":3000"))
}
