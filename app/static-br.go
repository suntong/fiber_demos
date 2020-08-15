package main

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
)

func main() {
	app := fiber.New()

	//app.Static("/", "./public")
	// => http://localhost:3000/js/script.js
	// => http://localhost:3000/css/style.css

	app.Static("/sample0", "../web/sample0")

	app.Static("/sample1", "../web/sample1")
	// => http://localhost:3000/sample1/js/script.js
	// => http://localhost:3000/sample1/css/style.css
	app.Static("/ot", "../web/OverTime")

	// Default compression config
	app.Use(middleware.Compress())
	// Register static route
	app.Static("/sample-br0", "../web/sample0")
	app.Static("/sample-br1", "../web/br/sample-br1")
	app.Static("/sample-br2", "../web/br/sample-br2")

	app.Listen(3000)
}
