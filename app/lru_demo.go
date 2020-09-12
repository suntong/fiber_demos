package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber"
	"github.com/gofiber/utils"
	"github.com/hashicorp/golang-lru/simplelru"
)

func main() {
	app := fiber.New()

	evictCounter := 0
	onEvicted := func(k interface{}, v interface{}) {
		evictCounter++
	}
	// https://godoc.org/github.com/hashicorp/golang-lru/simplelru
	l, err := simplelru.NewLRU(12, onEvicted)
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	// http://localhost:3000/q?p=something
	app.Get("/q", func(c *fiber.Ctx) {
		p := utils.ImmutableString(c.Query("p"))
		l.Add(p, nil)
		r := fmt.Sprintln(l.Keys())
		c.Send(r)
	})

	app.Listen(3000)
}

/*

https://github.com/gofiber/fiber/issues/763

$ curl http://localhost:3000/q?p=AAAAAAAAAAAAA
[AAAAAAAAAAAAA]

$ curl http://localhost:3000/q?p=BBBBBBBBBBBBB
[AAAAAAAAAAAAA BBBBBBBBBBBBB]

$ curl http://localhost:3000/q?p=CCCCCCCCCCCCC
[AAAAAAAAAAAAA BBBBBBBBBBBBB CCCCCCCCCCCCC]

$ curl http://localhost:3000/q?p=DDDDDDDDDDDDD
[AAAAAAAAAAAAA BBBBBBBBBBBBB CCCCCCCCCCCCC DDDDDDDDDDDDD]

$ curl http://localhost:3000/q?p=EEEEEEEEEEEEE
[AAAAAAAAAAAAA BBBBBBBBBBBBB CCCCCCCCCCCCC DDDDDDDDDDDDD EEEEEEEEEEEEE]

$ curl http://localhost:3000/q?p=CCCCCCCCCCCCC
[AAAAAAAAAAAAA BBBBBBBBBBBBB DDDDDDDDDDDDD EEEEEEEEEEEEE CCCCCCCCCCCCC]

*/
