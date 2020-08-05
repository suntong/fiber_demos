package main

import (
	"log"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) {
		c.Redirect("http://localhost:3000/hello")
	})
	app.Get("/hello", func(c *fiber.Ctx) {
		c.Send("Hello 👋!")
	})
	app.Get("/h", func(c *fiber.Ctx) {
		c.Redirect("https://github.com:443/gofiber/fiber")
	})
	app.Get("/g", func(c *fiber.Ctx) {
		c.Redirect("https://140.82.114.4:443/gofiber/fiber")
	})
	log.Fatal(app.Listen(3000))
}
