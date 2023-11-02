package znet

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/contrib/websocket"
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
	if s.g.NombreJoueurs < 4 {
		p := player.NewPlayer()
		p.Conn = c
		err := s.g.AddPlayer(p)
		log.Println(p, err)

		r := game.ReceiveTakeHandEvt(*p, gametake.AllTakes)
		m, err := json.Marshal(r)
		fmt.Println(err)

		if err := c.WriteMessage(1, m); err != nil {
			log.Println("write:", err)
		}
	}
}

func (s *SocketHandler) Handle(c *websocket.Conn) {
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
		_ = json.Unmarshal(msg, &obj)

		id, _ := strconv.Atoi(obj["id"])
		if id == 2 {
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
		fmt.Println("take not found", obj)
		return
	}

	err := s.g.AddTake(pid, tk)

	if err != nil {
		log.Println(err.Error())
	}

	log.Println(s.g.GetTake().Name())

	if !s.g.TakesFinished {
		b := game.BroadcastPlayerTakeEvt(obj["gametake"], id, s.g.AvailableTakes())

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
}

func (s *SocketHandler) HandlePlayerCard(_ *websocket.Conn, obj map[string]string) {
	pid, _ := strconv.Atoi(obj["player_id"])
	card := cards.Card{Couleur: obj["color"], Genre: obj["genre"]}

	fmt.Println(pid)

	err := s.g.PlayCard(pid, card)
	if err != nil {
		log.Println(err.Error())
	}

	deck, _ := s.g.CurrentDeck()

	for _, p := range s.g.GetPlayers() {
		if p != nil && p.Conn != nil {
			b := game.ReceiveDeckEvt(*p, deck)
			m, err := json.Marshal(b)
			fmt.Println(err)

			if err := p.Conn.WriteMessage(1, m); err != nil {
				log.Println("write:", err)
			}
		}
	}
}
