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

	// to serve directly from already-compressed files with the .br suffix
	app.Settings.CompressedFileSuffix = ".br" // default: ".fiber.gz"
	// https://docs.gofiber.io/api/app#settings
	// https://docs.gofiber.io/application#static
	app.Static("/sample-br1", "../web/sample-br", fiber.Static{
		Compress: true,
	})

	// See https://github.com/gofiber/fiber/issues/726
	// Static files can only be pre-compressed with GZip, not Brotli
	app.Settings.CompressedFileSuffix = ".gz" // default: ".fiber.gz"
	// https://docs.gofiber.io/api/app#settings
	// https://docs.gofiber.io/application#static
	app.Static("/sample-gz", "../web/sample-gz", fiber.Static{
		Compress: true,
	})

	app.Listen(3000)
}
