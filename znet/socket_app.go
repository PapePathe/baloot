package znet

import (
	"fmt"
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type sApp struct {
	app  *fiber.App
	port string
}

func NewSocketApp(port string) sApp {
	app := fiber.New()
	return sApp{
		app:  app,
		port: port,
	}
}

func (s sApp) Start() {
	log.Fatal(s.app.Listen(fmt.Sprintf(":%s", s.port)))
}

func (s sApp) Stop() {
	s.app.Shutdown()
}

func (s sApp) SetupRoutes(h SocketHandler) {
	s.app.Use("/ws", h.Upgrade)
	s.app.Get("/ws/:id", websocket.New(h.Handle))
}
