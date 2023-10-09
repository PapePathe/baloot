package player

import (
	"pathe.co/zinx/gametake"
	"pathe.co/zinx/pkg/cards"
)

type IHand interface {
	Play() cards.Card
	OrderCards(t *gametake.GameTake)
}
