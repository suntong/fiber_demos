package main

import "github.com/gofiber/fiber"

func main() {
	app := fiber.New()

	//app.Static("/", "./public")
	// => http://localhost:3000/js/script.js
	// => http://localhost:3000/css/style.css

	app.Static("/sample0", "../web/sample0")

	app.Static("/sample1", "../web/sample1")
	// => http://localhost:3000/sample1/js/script.js
	// => http://localhost:3000/sample1/css/style.css

	// to serve directly from already-compressed files with the .fiber.gz suffix
	app.Static("/sample2", "../web/sample2", fiber.Static{
		Compress: true,
	})

	app.Listen(3000)
}
