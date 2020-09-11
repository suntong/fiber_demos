package main

import (
	"fmt"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	// GET /dictionary.txt
	app.Get("/:file.:ext", func(c *fiber.Ctx) {
		msg := fmt.Sprintf("📃 %s.%s", c.Params("file"), c.Params("ext"))
		c.Send(msg) // => 📃 dictionary.txt
	})

	// GET /flights/LAX-SFO
	app.Get("/flights/:from-:to", func(c *fiber.Ctx) {
		msg := fmt.Sprintf("💸 From: %s, To: %s", c.Params("from"), c.Params("to"))
		c.Send(msg) // => 💸 From: LAX, To: SFO
	})

	// GET /api/register
	app.Get("/api/*", func(c *fiber.Ctx) {
		msg := fmt.Sprintf("✋ %s", c.Params("*"))
		c.Send(msg) // => ✋ /api/register
	})

	// GET /john
	app.Get("/:name", func(c *fiber.Ctx) {
		msg := fmt.Sprintf("Hello, %s 👋!", c.Params("name"))
		c.Send(msg) // => Hello john 👋!
	})

	// GET /john/75
	app.Get("/:name/:age/:gender?", func(c *fiber.Ctx) {
		msg := fmt.Sprintf("👴 %s is %s years old", c.Params("name"), c.Params("age"))
		c.Send(msg) // => 👴 john is 75 years old
	})

	app.Listen(3000)
}
