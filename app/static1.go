package main

import "github.com/gofiber/fiber"

func main() {
	app := fiber.New()

	//app.Static("/", "./public")
	// => http://localhost:3000/js/script.js
	// => http://localhost:3000/css/style.css

	app.Static("/sample0", "../web/sample0")
	// => http://localhost:3000/prefix/js/script.js
	// => http://localhost:3000/prefix/css/style.css

	app.Listen(3000)
}
