package player

import (
	"errors"

	"pathe.co/zinx/gametake"
	"pathe.co/zinx/pkg/cards"

	"github.com/rs/zerolog/log"
)

type machineDeck struct {
	cards    [4]cards.Card
	gametake gametake.GameTake
	hand     PlayingHand
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

func (m machineDeck) WinningOrLowestCard() cards.Card {
	found, pCards := m.hand.HasColor(m.cards[0].Couleur)

	if found {
		c, err := m.AttemptWin(pCards)

		if err != nil {
			log.Err(err).Msg("Current hand cannot win this deck going to play lowest card")
			return m.lowestCardOfColor(m.cards[0].Couleur, pCards)
		}

		log.Debug().Str("Card", c.String()).Msg("Going to play card of same color as in Deck")
		return c

	} else {
		for _, pc := range m.hand.Cards {
			if pc.Couleur != "" && pc.Genre != "" {
				log.Debug().Str("Card", pc.String()).Msg("Going to play card of other color because missing color in hand")
				return pc
			}
		}
	}

	return cards.Card{}
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

func (m machineDeck) LowestCard() cards.Card {
	return m.hand.LowestCard(m.gametake)
}

func (m machineDeck) lowestCardOfColor(color string, pCards []cards.Card) cards.Card {
	c := pCards[0]

	for _, cc := range pCards {
		if m.gametake.EvaluateCardForWin(c) > m.gametake.EvaluateCardForWin(cc) {
			c = cc
		}
	}

	return c
}
