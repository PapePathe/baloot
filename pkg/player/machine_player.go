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
		log.Info().Int("Machine Player ID", p.ID).Interface("Deck", e.Deck).Str("Hand", p.PlayingHand.String()).Msg("It's my turn to play")

		c := cards.Card{}

		if e.Deck[0].Couleur != "" {
			found, pCards := p.HasColor(e.Deck[0].Couleur)

			if found {
				c = pCards[0]

				log.Info().Int("Machine Player ID", p.ID).Str("Card", c.String()).Msg("Going to play card of same color as in Deck")
			} else {

				for _, pc := range p.PlayingHand.Cards {
					if pc.Couleur != "" && pc.Genre != "" {
						log.Info().Int("Machine Player ID", p.ID).Str("Card", pc.String()).Msg("Going to play card of other color because missing color in hand")
						c = pc
					}
				}

			}
		} else {
			for _, pc := range p.PlayingHand.Cards {
				if pc.Couleur != "" && pc.Genre != "" {
					log.Info().Int("Machine Player ID", p.ID).Str("Card", pc.String()).Msg("Going to play the first card on deck")
					c = pc
				}
			}
		}

		e.PlayChannel <- PlayEvent{Card: c, PlayerID: p.GetID()}

		time.Sleep(time.Second)
	}
}

func (p *MachinePlayer) BroadCastPlayerTake(e BroadcastPlayerTakeMsg) {
	fmt.Println("Received player take", e)
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
	fmt.Println("Received hand", h)
	p.Hand = h
}

func (p *MachinePlayer) SetPlayingHand(h PlayingHand) {
	fmt.Println("Received playing hand", h)
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

func (p *MachinePlayer) HasColor(c string) (bool, []cards.Card) {
	foundCards := []cards.Card{}

	for _, pc := range p.PlayingHand.Cards {
		if pc.Couleur == c {
			foundCards = append(foundCards, pc)
		}
	}

	return len(foundCards) > 0, foundCards
}

func (p *MachinePlayer) HasCard(c cards.Card) (bool, int) {
	for idx, pc := range p.PlayingHand.Cards {
		if pc == c {
			return true, idx
		}
	}

	return false, -1
}
