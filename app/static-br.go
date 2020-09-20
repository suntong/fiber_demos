package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

func main() {
	app := fiber.New()

	// http://127.0.0.1:3000/sample0/
	// Content-Type: text/html; charset=utf-8
	app.Static("/sample0", "../web/sample0")
	app.Static("/sample1", "../web/sample1")
	app.Static("/ot", "../web/OverTime")

	app.Use(compress.New())

	app.Get("/", func(c *fiber.Ctx) error {
		c.Request().Header.Set(fiber.HeaderAcceptEncoding, "br")
		return c.Next()
	})

	// http://127.0.0.1:3000/sample-br0/
	// Content-Encoding: br
	app.Static("/sample-br0", "../web/sample0")
	app.Static("/sample-br1", "../web/br/sample-br1")
	app.Static("/sample-br2", "../web/br/sample-br2")

	app.Listen(":3000")
}
