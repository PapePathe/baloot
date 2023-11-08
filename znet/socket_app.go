package znet

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/template/html/v2"
)

type SocketApp struct {
	app  *fiber.App
	port string
}

//go:embed views/*
var viewsfs embed.FS

func NewSocketApp(port string) SocketApp {
	engine := html.NewFileSystem(http.FS(viewsfs), ".html")
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "views/layouts/main",
	})

	app.Static("/assets", "./public", fiber.Static{
		CacheDuration: 10 * time.Second,
		Compress:      true,
		MaxAge:        3600,
	})
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format:     "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Africa/Dakar",
	}))

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
		log.Println(fmt.Errorf(" error shutting down app %w", err))
	}
}

func (s SocketApp) SetupRoutes(h SocketHandler) {
	s.app.Use("/ws", h.Upgrade)
	s.app.Get("/", h.Index)
	s.app.Get("/ws/:id", websocket.New(h.Handle))
}
