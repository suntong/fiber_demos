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

	// to serve directly from already-compressed files with the .gz suffix
	app.Settings.CompressedFileSuffix = ".gz" // default: ".fiber.gz"
	// https://docs.gofiber.io/api/app#settings
	// https://docs.gofiber.io/application#static
	app.Static("/sample-gz", "../web/sample-gz", fiber.Static{
		Compress: true,
	})

	app.Listen(3000)
}
