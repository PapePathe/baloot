package player

import (
	"fmt"
	"strings"

	"pathe.co/zinx/gametake"
)

type IPlayerTakes interface {
	GetBestTake() gametake.GameTake
	GetTakes() map[gametake.GameTake]gametake.GameTakeEntry
	SetTake(*gametake.GameTake)
}

func (p *Player) GetTakes() map[gametake.GameTake]gametake.GameTakeEntry {
	m := make(map[gametake.GameTake]gametake.GameTakeEntry)

	m[&gametake.CARREAU] = gametake.CARREAU.EvaluateHand(p.Hand.Cards)
	m[gametake.CENT] = gametake.CENT.EvaluateHand(p.Hand.Cards)
	m[gametake.COEUR] = gametake.COEUR.EvaluateHand(p.Hand.Cards)
	m[gametake.PIQUE] = gametake.PIQUE.EvaluateHand(p.Hand.Cards)
	m[gametake.TREFLE] = gametake.TREFLE.EvaluateHand(p.Hand.Cards)
	m[gametake.TOUT] = gametake.TOUT.EvaluateHand(p.Hand.Cards)

	return m
}

func (p *Player) ShowTakes() string {
	var sb strings.Builder
	for key, element := range p.GetTakes() {
		sb.WriteString(fmt.Sprintf("(%v, %v)", key.Name(), element))
	}

	return sb.String()
}

func (p *Player) GetBestTake() (take gametake.GameTake) {

	takes := []gametake.GameTake{}

	for take, takeEntry := range p.GetTakes() {
		if takeEntry.CanTake(take) {
			takes = append(takes, take)
		}
	}

	if len(takes) == 1 {
		take = takes[0]
		return take
	}

	take = gametake.PASSE

	return take
}

func (p *Player) SetTake(gt *gametake.GameTake) {
	p.Take = gt
}
