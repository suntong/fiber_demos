package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

const cookieName = "test_cookie"

func main() {
	app := fiber.New()

	// http://localhost:3000/q?p=something
	app.Get("/q", func(c *fiber.Ctx) error {
		cookie := new(fiber.Cookie)
		cookie.Expires = time.Now().Add(24 * time.Hour)
		tests := c.Cookies(cookieName)
		tests += ", " + c.Query("p")

		// Set cookie
		cookie.Name = cookieName
		cookie.Value = tests
		c.Cookie(cookie)

		// Get cookie value right back
		testn := c.Cookies(cookieName)
		return c.SendString(tests + " => " + testn)
	})

	// get cookie (this should be ran after the above)
	app.Get("/c", func(c *fiber.Ctx) error {
		return c.SendStringString(c.Cookies(cookieName))
	})

	app.Listen(":3000")
}

/*

See https://github.com/gofiber/fiber/issues/768

Repeatably make request to http://localhost:3000/q?p=AAA will get

, AAA, AAA, AAA, AAA => , AAA, AAA, AAA

This is normal as

- when setting the cookie with c.Cookie, it adds the Set-Cookie header to the response object
- I.e., Cookies are set in the response header, and retrieved by looking at the request header.

*/
