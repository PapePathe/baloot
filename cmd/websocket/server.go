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

			r := ReceiveTakeHandMsg(*p, gametake.AllTakes)
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
				log.Println("read:", err)
				break
			}

			log.Printf("recv: %s %s", msg, mt)
			obj := map[string]string{}
			_ = json.Unmarshal(msg, &obj)

			id, _ := strconv.Atoi(obj["id"])
			if id == 2 {
				pid, _ := strconv.Atoi(obj["player_id"])
				err := g.AddTake(pid, gametake.AllTakesByName[obj["gametake"]])
				if err != nil {
					log.Printf(err.Error())
				}

				log.Printf(g.GetTake().Name())

				b := broadcastPlayerTake{ID: BroadcastPlayerTake, Take: obj["gametake"], PlayerId: id, AvailableTakes: g.AvailableTakes()}
				for _, p := range g.GetPlayers() {
					m, _ := json.Marshal(b)
					if err := p.Conn.WriteMessage(1, m); err != nil {
						log.Println("write:", err)
					}
				}
			}
		}
	}))

	log.Fatal(app.Listen(":7777"))
}

type messageID int

var (
	ReceiveTakeHand     messageID = 1
	ReceivePlayingHand  messageID = 2
	SetTake             messageID = 3
	PlayCard            messageID = 4
	BroadcastCards      messageID = 5
	BroadcastPlayerTake messageID = 5
)

type broadcastPlayerTake struct {
	ID             messageID           `json:"id"`
	Take           string              `json: 'take'`
	PlayerId       int                 `json:"player_id"`
	AvailableTakes []gametake.GameTake `json:"available_takes"`
}

type receiveTakeHandMsg struct {
	ID             messageID           `json:"id"`
	Player         player.Player       `json:"player"`
	AvailableTakes []gametake.GameTake `json:"available_takes"`
}

func ReceiveTakeHandMsg(p player.Player, takes []gametake.GameTake) receiveTakeHandMsg {
	clientMessage := receiveTakeHandMsg{ID: ReceiveTakeHand, Player: p, AvailableTakes: takes}

	return clientMessage
}

type setTake struct {
	ID       messageID `json:"id"`
	PlayerId int       `json:"player_id"`
	Gametake string    `json:"gametake"`
}

func setTakeMsg(gt string, pid int) setTake {
	clientMessage := setTake{ID: SetTake, PlayerId: pid, Gametake: gt}

	return clientMessage
}
