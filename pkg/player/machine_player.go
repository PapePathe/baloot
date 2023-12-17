package player

import (
	"fmt"
	"time"

	"os"

	"github.com/gofiber/contrib/websocket"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"pathe.co/zinx/gametake"
	"pathe.co/zinx/pkg/cards"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

type MachinePlayer struct {
	Hand        Hand               `json:"hand"`
	PlayingHand PlayingHand        `json:"playingHand"`
	Take        *gametake.GameTake `json:"take"`
	ID          int                `json:"id"`
}

func NewMachinePlayer() *MachinePlayer {
	return &MachinePlayer{
		Hand:        Hand{Cards: [5]cards.Card{}},
		PlayingHand: PlayingHand{Cards: []cards.Card{}},
		Take:        nil,
		ID:          0,
	}
}

func (p *MachinePlayer) BroadCastGameDeck(e ReceiveDeckMsg) {
	if e.NextPlayer == p.ID {
		time.Sleep(time.Second)
		log.Debug().Int("Machine Player ID", p.ID).Str("Hand", p.PlayingHand.String()).Msg("It's my turn to play")
		log.Trace().Msg(fmt.Sprintf("Deck: %s %s %s %s", e.Deck[0], e.Deck[1], e.Deck[2], e.Deck[3]))

		c := cards.Card{}
		mDeck := machineDeck{cards: e.Deck, gametake: e.Take, hand: p.PlayingHand}

		if e.Deck[0].Couleur != "" {
			log.Debug().Int("Machine Player ID", p.ID).Msg("==> Searching for winning or lowest card")
			c = mDeck.WinningOrLowestCard()
			log.Debug().Int("Machine Player ID", p.ID).Msg("==> Found card")
		} else {
			c = mDeck.LowestCard()
			log.Debug().Int("Machine Player ID", p.ID).Str("Card", c.String()).Msg("Going to play the lowest card since player does not have card")
		}

		if !c.IsNotEmpty() {
			panic(c)
		}
		e.PlayChannel <- PlayEvent{Card: c, PlayerID: p.GetID()}

		time.Sleep(time.Second)
	}
}

func (p *MachinePlayer) BroadCastPlayerTake(e BroadcastPlayerTakeMsg) {
	log.Debug().Str("Received player take", e.Take)
}

func (p *MachinePlayer) SetTake(t *gametake.GameTake) {
	p.Take = t
}

func (p *MachinePlayer) GetTake() *gametake.GameTake {
	return p.Take
}

func (p *MachinePlayer) SetID(id int) {
	p.ID = id
}

func (p *MachinePlayer) SetHand(h Hand) {
	log.Trace().Msg(fmt.Sprintf("Received hand %s", h))
	p.Hand = h
}

func (p *MachinePlayer) SetPlayingHand(h PlayingHand) {
	log.Trace().Int("Machine player", p.ID).Msg(fmt.Sprintf("Received playing hand %s", h))
	p.PlayingHand = h
}

func (p *MachinePlayer) GetPlayingHand() PlayingHand {
	return p.PlayingHand
}

func (p *MachinePlayer) GetHand() Hand {
	return p.Hand
}

func (p *MachinePlayer) GetConn() *websocket.Conn {
	return nil
}

func (p *MachinePlayer) GetID() int {
	return p.ID
}

func (p *MachinePlayer) HasCard(c cards.Card) (bool, int) {
	for idx, pc := range p.PlayingHand.Cards {
		if pc == c {
			return true, idx
		}
	}

	return false, -1
}
