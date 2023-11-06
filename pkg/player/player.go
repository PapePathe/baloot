package player

import (
	"sort"

	"github.com/gofiber/contrib/websocket"
	"pathe.co/zinx/gametake"
	"pathe.co/zinx/pkg/cards"
)

type Player struct {
	Hand        Hand               `json:"hand"`
	PlayingHand PlayingHand        `json:"playingHand"`
	Take        *gametake.GameTake `json:"take"`
	ID          int                `json:"id"`
	Conn        *websocket.Conn    `json:"-"`
}

func NewPlayer() *Player {
	return &Player{
		Hand:        Hand{Cards: [5]cards.Card{}},
		PlayingHand: PlayingHand{Cards: []cards.Card{}},
		Take:        nil,
		ID:          0,
		Conn:        nil,
	}
}

func (p *Player) SetID(id int) {
	p.ID = id
}

func (p *Player) GetID() int {
	return p.ID
}

func (p *Player) HasCard(c cards.Card) (bool, int) {
	for idx, pc := range p.PlayingHand.Cards {
		if pc == c {
			return true, idx
		}
	}

	return false, -1
}

func (p *Player) OrderedCards() map[string][]cards.Card {
	cardsMap := make(map[string][]cards.Card)

	for _, card := range p.Hand.Cards {
		_, ok := cardsMap[card.Couleur]

		if ok {
			cardsMap[card.Couleur] = append(cardsMap[card.Couleur], card)
		} else {
			cardsMap[card.Couleur] = []cards.Card{card}
		}
	}

	return cardsMap
}

func (p *Player) OrderedCardsForTake(take gametake.GameTake) [5]cards.Card {
	cardsMap := make(map[string][]cards.Card)

	for _, card := range p.Hand.Cards {
		_, ok := cardsMap[card.Couleur]

		if ok {
			cardsMap[card.Couleur] = append(cardsMap[card.Couleur], card)
		} else {
			cardsMap[card.Couleur] = []cards.Card{card}
		}

		sorter := SortByColorAndType{cardsMap[card.Couleur], take}
		sort.Sort(sorter)
	}

	result := []cards.Card{}

	for _, cards := range cardsMap {
		result = append(result, cards...)
	}

	return [5]cards.Card(result)
}

func (p *Player) OrderedCardsForPlaying(take gametake.GameTake) [8]cards.Card {
	cardsMap := make(map[string][]cards.Card)

	for _, card := range p.PlayingHand.Cards {
		_, ok := cardsMap[card.Couleur]

		if ok {
			cardsMap[card.Couleur] = append(cardsMap[card.Couleur], card)
		} else {
			cardsMap[card.Couleur] = []cards.Card{card}
		}

		sorter := SortByColorAndType{cardsMap[card.Couleur], take}
		sort.Sort(sorter)
	}

	result := []cards.Card{}

	for _, cards := range cardsMap {
		result = append(result, cards...)
	}

	return [8]cards.Card(result)
}
