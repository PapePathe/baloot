package znet

import (
	"context"
	"encoding/json"
	"os"
	"strconv"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	gt_proto "pathe.co/zinx/proto/gametake_history/v1"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"pathe.co/zinx/gametake"
	"pathe.co/zinx/pkg/cards"
	"pathe.co/zinx/pkg/game"
	"pathe.co/zinx/pkg/player"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

type SocketHandler struct {
	g                  *game.Game
	gTakeHistoryClient gt_proto.GameTakeHistoryClient
}

func NewSocketHandler() SocketHandler {
	takesSvcAddr := "localhost:50052"
	if os.Getenv("ZINX_TAKES_SERVER") != "" {
		takesSvcAddr = os.Getenv("ZINX_TAKES_SERVER")
	}

	conn, err := grpc.Dial(takesSvcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error().Err(err).Msg("did not connect to gametake history service")
	}

	cli := gt_proto.NewGameTakeHistoryClient(conn)

	return SocketHandler{g: game.NewGame(), gTakeHistoryClient: cli}
}

func (s *SocketHandler) StartPlayerRegistration(c *websocket.Conn) {
	log.Debug().Msg("starting player registration")

	if s.g.NombreJoueurs < 4 {
		log.Debug().Msg("adding new player")

		p := player.NewPlayer()
		p.Conn = c
		err := s.g.AddPlayer(p)

		if err != nil {
			log.Warn().Err(err).Msg("error adding player : ")
		}

		r := game.ReceiveTakeHandEvt(*p, gametake.AllTakeNames)
		m, err := json.Marshal(r)

		if err != nil {
			log.Error().Err(err).Msg("marshaling take hand msg")
		}

		if err := c.WriteMessage(1, m); err != nil {
			log.Error().Err(err).Msg("error writing message to socket")
		}

		s.g.AddPlayer(player.NewMachinePlayer())
		s.g.AddPlayer(player.NewMachinePlayer())
		s.g.AddPlayer(player.NewMachinePlayer())

		go s.g.StartPlayChannel()
	}
}

func (s *SocketHandler) Index(c *fiber.Ctx) error {
	return c.Render("views/index", fiber.Map{
		"Title": "Hello, Your are in the playground!",
	})
}

func (s *SocketHandler) Upgrade(c *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) {
		log.Debug().Msg("Socket neeeds upgrade")
		c.Locals("allowed", true)

		return c.Next()
	}

	return fiber.ErrUpgradeRequired
}

func (s *SocketHandler) Handle(c *websocket.Conn) {
	log.Debug().Msg("Start new connection handler")
	s.StartPlayerRegistration(c)

	var (
		mt  int
		msg []byte
		err error
	)

	for {
		if mt, msg, err = c.ReadMessage(); err != nil {
			log.Error().Err(err).Msg("error reading message:")

			break
		}

		log.Debug().Str("Message", string(msg)).Int("MessageType", mt).Msg("")

		obj := map[string]string{}
		err = json.Unmarshal(msg, &obj)

		if err != nil {
			log.Error().Err(err).Msg("")

			break
		}

		id, _ := strconv.Atoi(obj["id"])

		if id == 2 {
			log.Debug().Msg("handling player take")
			s.HandlePlayerTake(c, obj)
		}

		if id == 4 {
			s.HandlePlayerCard(c, obj)
		}
	}
}

func (s *SocketHandler) saveTakeHistory(pid int) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	cards := []*gt_proto.Card{}
	for _, c := range s.g.GetPlayers()[pid].GetHand().Cards {
		cards = append(cards, &gt_proto.Card{Type: c.Genre, Color: c.Couleur})
	}

	playerTakes := []string{}
	for _, p := range s.g.GetPlayers() {
		if p != nil && p.GetID() != pid && p.GetTake() != nil {
			playerTakes = append(playerTakes, (*p.GetTake()).Name())
		}
	}
	resp, err := s.gTakeHistoryClient.Add(ctx, &gt_proto.GameTakeHistoryRequest{
		Constraints: playerTakes,
		Take:        (*s.g.GetPlayers()[pid].GetTake()).Name(),
		Cards:       cards,
	})

	if err != nil {
		return err
	}

	log.Info().Str("Save take history response", resp.String())

	return nil
}

func (s *SocketHandler) HandlePlayerTake(_ *websocket.Conn, obj map[string]string) {
	id, _ := strconv.Atoi(obj["id"])
	pid, _ := strconv.Atoi(obj["player_id"])
	tk, ok := gametake.AllTakesByName[obj["gametake"]]

	if !ok {
		log.Error().Str("GameTake", obj["gametake"]).Msg("could not find gametake")

		return
	}

	err := s.g.AddTake(pid, tk)

	if err != nil {
		log.Error().Err(err).Msg("error adding take")

		return
	}

	err = s.saveTakeHistory(pid)

	if err != nil {
		log.Error().Err(err).Caller().Msg("error saving take history")
	}

	if !s.g.TakesFinished {
		log.Trace().Msg("Takes are not finished yet")
	}

	s.broadcastPlayerTake(id, obj, tk)
}

func (s *SocketHandler) broadcastPlayerTake(id int, obj map[string]string, tk gametake.GameTake) {
	takes := []string{}

	for _, t := range s.g.AvailableTakes() {
		if t == gametake.PASSE || t.GreaterThan(tk) {
			takes = append(takes, t.Name())
		}
	}

	b := player.BroadcastPlayerTakeEvt(obj["gametake"], id, takes)

	for _, p := range s.g.GetPlayers() {
		if p != nil {
			p.BroadCastPlayerTake(b)
		}
	}
}

func (s *SocketHandler) HandlePlayerCard(_ *websocket.Conn, obj map[string]string) {
	pid, _ := strconv.Atoi(obj["player_id"])
	card := cards.Card{Couleur: obj["color"], Genre: obj["genre"]}

	err := s.g.PlayCardNext(pid, card)
	if err != nil {
		log.Error().Err(err).Msg("error playing card")
	}
}
