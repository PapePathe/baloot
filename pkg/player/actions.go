package player

import "pathe.co/zinx/pkg/cards"

type IPlayerActions interface {
	PlayCard(c cards.Card)
}
