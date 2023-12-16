package player

import (
	"errors"
	"pathe.co/zinx/gametake"
	"pathe.co/zinx/pkg/cards"
)

type machineDeck struct {
	cards    [4]cards.Card
	gametake gametake.GameTake
}

func (m machineDeck) FindWinner() cards.Card {
	winningCard := m.cards[0]

	for _, c := range m.cards {
		if c.IsNotEmpty() && winningCard.IsNotEmpty() {
			winner := m.gametake.Winner(winningCard, c)

			if winner == c {
				winningCard = c
			}
		}
	}

	return winningCard
}

func (m machineDeck) AttemptWin(hand []cards.Card) (cards.Card, error) {
	w := m.FindWinner()

	for _, c := range hand {
		if c.IsNotEmpty() && w.IsNotEmpty() && m.gametake.Winner(w, c) == c {

			return c, nil
		}
	}

	return cards.Card{}, errors.New("Could not find card that can win")
}
