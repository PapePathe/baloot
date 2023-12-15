package player

import (
	"fmt"
	"sort"

	"github.com/gofiber/contrib/websocket"
	"pathe.co/zinx/gametake"
	"pathe.co/zinx/pkg/cards"
)

type BelotePlayer interface {
	GetID() int
	SetID(i int)
	HasCard(cards.Card) (bool, int)
	SetTake(*gametake.GameTake)
	GetTake() *gametake.GameTake
	GetPlayingHand() PlayingHand
	GetHand() Hand
	SetPlayingHand(PlayingHand)
	GetConn() *websocket.Conn
}

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

func (p *Player) GetTake() *gametake.GameTake {
	return p.Take
}
func (p *Player) SetID(id int) {
	p.ID = id
}

func (p *Player) SetPlayingHand(h PlayingHand) {
	p.PlayingHand = h
}

func (p *Player) GetPlayingHand() PlayingHand {
	return p.PlayingHand
}

func (p *Player) GetHand() Hand {
	return p.Hand
}

func (p *Player) GetConn() *websocket.Conn {
	return p.Conn
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

	keys := make([]string, 0, len(cardsMap))
	for k := range cardsMap {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	fmt.Println(keys)

	result := []cards.Card{}

	for _, key := range keys {
		mapCards := cardsMap[key]
		result = append(result, mapCards...)
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

	keys := make([]string, 0, len(cardsMap))
	for k := range cardsMap {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	fmt.Println(keys)

	result := []cards.Card{}

	for _, key := range keys {
		mapCards := cardsMap[key]
		result = append(result, mapCards...)
	}

	return [8]cards.Card(result)
}
