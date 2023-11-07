package znet

import (
	"fmt"
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type SocketApp struct {
	app  *fiber.App
	port string
}

func NewSocketApp(port string) SocketApp {
	app := fiber.New()

	return SocketApp{
		app:  app,
		port: port,
	}
}

func (s SocketApp) Start() {
	log.Fatal(s.app.Listen(fmt.Sprintf(":%s", s.port)))
}

func (s SocketApp) Stop() {
	err := s.app.Shutdown()
	if err != nil {
		log.Println(err)
	}
}

func (s SocketApp) SetupRoutes(h SocketHandler) {
	s.app.Use("/ws", h.Upgrade)
	s.app.Get("/ws/:id", websocket.New(h.Handle))
}
