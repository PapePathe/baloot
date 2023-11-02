package game

import (
	"errors"

	"pathe.co/zinx/gametake"
	"pathe.co/zinx/pkg/cards"
)

var (
	ErrCannotAddMoreThanFourCardsToDeck = errors.New("cannot add more than five cards to deck")
	ErrCannotAddExistingCardToDeck      = errors.New("cannot add existing card to deck")
	ErrCannotAddEmptyCardToDeck         = errors.New("cannot add empty card to deck")
	ErrNotYourTurnToPlay                = errors.New("cannot add card, it is not your turn to play")
)

type Deck struct {
	cards      [4]cards.Card
	cardscount int
	players    [4]int
	winner     int
	gametake   gametake.GameTake
}

func NewDeck(p [4]int, gt gametake.GameTake) Deck {
	return Deck{
		winner:     -1,
		players:    p,
		gametake:   gt,
		cards:      [4]cards.Card{},
		cardscount: 0,
	}
}

func (d *Deck) AddCard(pid int, card cards.Card) error {
	if err := d.validateCard(card); err != nil {
		return err
	}

	if pid != d.cardscount {
		return ErrNotYourTurnToPlay
	}

	d.cards[d.cardscount] = card
	d.cardscount++

	if d.cardscount == 4 {
		d.findWinner()
	}

	return nil
}

func (d *Deck) Score() (int, int) {
	return 0, 0
}

func (d *Deck) validateCard(card cards.Card) error {
	if d.cardscount == 4 {
		return ErrCannotAddMoreThanFourCardsToDeck
	}

	if card.Genre == "" && card.Couleur == "" {
		return ErrCannotAddEmptyCardToDeck
	}

	if d.hasCard(card) {
		return ErrCannotAddExistingCardToDeck
	}

	return nil
}

func (d *Deck) hasCard(c cards.Card) bool {
	for _, dc := range d.cards {
		if dc == c {
			return true
		}
	}

	return false
}

func (d *Deck) findWinner() int {
	d.winner = d.players[0]
	winningCard := d.cards[0]

	for idx := 1; idx < 4; idx++ {
		currentCard := d.cards[idx]
		winner := d.gametake.Winner(currentCard, winningCard)

		if winner == currentCard {
			winningCard = currentCard
			d.winner = d.players[idx]
		}
	}

	return d.winner
}
