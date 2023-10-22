package main

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"pathe.co/zinx/gametake"
	"pathe.co/zinx/pkg/game"
	"pathe.co/zinx/pkg/player"
)

func main() {
	app := fiber.New()
	g := game.NewGame()

	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws/:id", websocket.New(func(c *websocket.Conn) {
		if g.NombreJoueurs < 4 {
			p := player.NewPlayer()
			p.Conn = c
			g.AddPlayer(p)
			log.Println(p)

			r := game.ReceiveTakeHandMsg(*p, gametake.AllTakes)
			m, _ := json.Marshal(r)
			if err := c.WriteMessage(1, m); err != nil {
				log.Println("write:", err)
			}
		}
		var (
			mt  int
			msg []byte
			err error
		)
		for {
			if mt, msg, err = c.ReadMessage(); err != nil {
				log.Println("error reading message:", err)
				break
			}

			log.Printf("recv: %s %d", msg, mt)
			obj := map[string]string{}
			_ = json.Unmarshal(msg, &obj)

			id, _ := strconv.Atoi(obj["id"])
			if id == 2 {
				pid, _ := strconv.Atoi(obj["player_id"])
				err := g.AddTake(pid, gametake.AllTakesByName[obj["gametake"]])
				if err != nil {
					log.Println(err.Error())
				}
				log.Println(g.GetTake().Name())

				b := game.BroadcastPlayerTakeMsg(obj["gametake"], id, g.AvailableTakes())
				for _, p := range g.GetPlayers() {
					if p != nil {
						m, _ := json.Marshal(b)
						if err := p.Conn.WriteMessage(1, m); err != nil {
							log.Println("write:", err)
						}
					}
				}
			}
		}
	}))

	log.Fatal(app.Listen(":7777"))
}
