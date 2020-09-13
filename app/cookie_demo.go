package main

import (
	"time"

	"github.com/gofiber/fiber"
)

const cookieName = "test_cookie"

func main() {
	app := fiber.New()

	// http://localhost:3000/q?p=something
	app.Get("/q", func(c *fiber.Ctx) {
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
		c.Send(tests + " => " + testn)
	})

	app.Listen(3000)
}

/*

$ curl http://localhost:3000/q?p=AAA
[AAAAAAAAAAAAA]

$ curl http://localhost:3000/q?p=BBB
[AAAAAAAAAAAAA BBBBBBBBBBBBB]

$ curl http://localhost:3000/q?p=CCC
[AAAAAAAAAAAAA BBBBBBBBBBBBB CCCCCCCCCCCCC]

$ curl http://localhost:3000/q?p=DDD
[AAAAAAAAAAAAA BBBBBBBBBBBBB CCCCCCCCCCCCC DDDDDDDDDDDDD]

$ curl http://localhost:3000/q?p=EEEEEEEEEEEEE
[AAAAAAAAAAAAA BBBBBBBBBBBBB CCCCCCCCCCCCC DDDDDDDDDDDDD EEEEEEEEEEEEE]

$ curl http://localhost:3000/q?p=CCCCCCCCCCCCC
[AAAAAAAAAAAAA BBBBBBBBBBBBB DDDDDDDDDDDDD EEEEEEEEEEEEE CCCCCCCCCCCCC]

*/
