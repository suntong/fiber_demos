package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	// GET http://localhost:3000/shoes?order=desc&brand=nike
	// GET http://localhost:3000/shoes?order=new&brand=adidas
	// GET http://localhost:3000/shoes?order=try&brand=joes&empty=a%2Fb%2Fc%2Cd
	app.Get("/shoes", func(c *fiber.Ctx) error {
		r := c.Query("order")                // "desc"
		r += "\n" + c.Query("brand")         // "nike"
		r += "\n" + c.Query("empty", "nike") // "nike"
		return c.SendString(r)
	})

	app.Listen(":3000")
}
