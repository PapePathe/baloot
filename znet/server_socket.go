package znet

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"pathe.co/zinx/gametake"
	"pathe.co/zinx/pkg/cards"
	"pathe.co/zinx/pkg/game"
	"pathe.co/zinx/pkg/player"
)

type SocketHandler struct {
	g *game.Game
}

func NewSocketHandler() SocketHandler {
	return SocketHandler{g: game.NewGame()}
}

func (s *SocketHandler) StartPlayerRegistration(c *websocket.Conn) {
	log.Println("starting player registration")

	if s.g.NombreJoueurs < 4 {
		log.Println("adding new player")

		p := player.NewPlayer()
		p.Conn = c
		err := s.g.AddPlayer(p)

		if err != nil {
			log.Println("error adding player : ", err)
		}

		r := game.ReceiveTakeHandEvt(*p, gametake.AllTakeNames)
		m, err := json.Marshal(r)

		if err != nil {
			log.Println("marshaling take hand msg", err)
		}

		if err := c.WriteMessage(1, m); err != nil {
			log.Println("write:", err)
		}
	}
}

func (s *SocketHandler) Upgrade(c *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) {
		fmt.Println("Socket neeeds upgrade")
		c.Locals("allowed", true)

		return c.Next()
	}

	return fiber.ErrUpgradeRequired
}

func (s *SocketHandler) Handle(c *websocket.Conn) {
	fmt.Println("Start new connection handler")
	s.StartPlayerRegistration(c)

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
		err = json.Unmarshal(msg, &obj)

		if err != nil {
			log.Println(err)

			break
		}

		log.Println(obj)

		id, _ := strconv.Atoi(obj["id"])

		if id == 2 {
			log.Println("handling player take")
			s.HandlePlayerTake(c, obj)
		}

		if id == 4 {
			s.HandlePlayerCard(c, obj)
		}
	}
}

func (s *SocketHandler) HandlePlayerTake(_ *websocket.Conn, obj map[string]string) {
	id, _ := strconv.Atoi(obj["id"])
	pid, _ := strconv.Atoi(obj["player_id"])
	tk, ok := gametake.AllTakesByName[obj["gametake"]]

	if !ok {
		log.Println("could not find gametake", obj["gametake"])

		return
	}

	err := s.g.AddTake(pid, tk)

	if err != nil {
		log.Println("error adding take", err.Error())
	}

	if !s.g.TakesFinished {
		log.Println("Takes are not finished yet")

		return
	}

	takes := []string{}

	for _, t := range s.g.AvailableTakes() {
		takes = append(takes, t.Name())
	}

	b := game.BroadcastPlayerTakeEvt(obj["gametake"], id, takes)

	for _, p := range s.g.GetPlayers() {
		if p != nil {
			m, err := json.Marshal(b)
			if err != nil {
				fmt.Println(err)
			}

			if err := p.Conn.WriteMessage(1, m); err != nil {
				log.Println("write:", err)
			}
		}
	}
}

func (s *SocketHandler) HandlePlayerCard(_ *websocket.Conn, obj map[string]string) {
	pid, _ := strconv.Atoi(obj["player_id"])
	card := cards.Card{Couleur: obj["color"], Genre: obj["genre"]}

	err := s.g.PlayCardNext(pid, card)
	if err != nil {
		log.Println(err.Error())
	}

	deck, _ := s.g.CurrentDeck()

	for _, p := range s.g.GetPlayers() {
		if p != nil && p.Conn != nil {
			scoreTeamA, scoreTeamB := s.g.Score()
			b := game.ReceiveDeckEvt(*p, deck, scoreTeamA, scoreTeamB)
			m, err := json.Marshal(b)
			fmt.Println(err)

			if err := p.Conn.WriteMessage(1, m); err != nil {
				log.Println("write:", err)
			}
		}
	}
}
