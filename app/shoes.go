package main

import "github.com/gofiber/fiber"

func main() {
	app := fiber.New()

	// GET http://localhost:3000/shoes?order=desc&brand=nike
	// GET http://localhost:3000/shoes?order=new&brand=adidas
	// GET http://localhost:3000/shoes?order=try&brand=joes&empty=a%2Fb%2Fc%2Cd
	app.Get("/shoes", func(c *fiber.Ctx) {
		r := c.Query("order")                // "desc"
		r += "\n" + c.Query("brand")         // "nike"
		r += "\n" + c.Query("empty", "nike") // "nike"
		c.Send(r)
	})

	app.Listen(3000)
}
