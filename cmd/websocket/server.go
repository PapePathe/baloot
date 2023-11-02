package main

import (
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"pathe.co/zinx/znet"
)

func main() {
	app := fiber.New()
	s := znet.NewSocketHandler()

	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)

			return c.Next()
		}

		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws/:id", websocket.New(s.Handle))

	log.Fatal(app.Listen(":7777"))
}
