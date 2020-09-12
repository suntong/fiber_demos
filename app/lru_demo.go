package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber"
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
		l.Add(c.Query("p"), nil)
		r := fmt.Sprintln(l.Keys())
		c.Send(r)
	})

	app.Listen(3000)
}
