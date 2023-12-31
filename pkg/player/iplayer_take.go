package player

import (
	"fmt"
	"strings"

	"pathe.co/zinx/gametake"
)

type IPlayerTakes interface {
	GetBestTake() gametake.GameTake
	GetTakes() map[gametake.GameTake]gametake.GameTakeEntry
	SetTake(take *gametake.GameTake)
}

func (p *Player) GetTakes() map[gametake.GameTake]gametake.GameTakeEntry {
	takes := make(map[gametake.GameTake]gametake.GameTakeEntry)

	takes[&gametake.CARREAU] = gametake.CARREAU.EvaluateHand(p.Hand.Cards)
	takes[gametake.CENT] = gametake.CENT.EvaluateHand(p.Hand.Cards)
	takes[gametake.COEUR] = gametake.COEUR.EvaluateHand(p.Hand.Cards)
	takes[gametake.PIQUE] = gametake.PIQUE.EvaluateHand(p.Hand.Cards)
	takes[gametake.TREFLE] = gametake.TREFLE.EvaluateHand(p.Hand.Cards)
	takes[gametake.TOUT] = gametake.TOUT.EvaluateHand(p.Hand.Cards)

	return takes
}

func (p *Player) ShowTakes() string {
	var sb strings.Builder
	for key, element := range p.GetTakes() {
		sb.WriteString(fmt.Sprintf("(%v, %v)", key.Name(), element))
	}

	return sb.String()
}

func (p *Player) GetBestTake() gametake.GameTake {
	takes := []gametake.GameTake{}

	for take, takeEntry := range p.GetTakes() {
		if takeEntry.CanTake(take) {
			takes = append(takes, take)
		}
	}

	if len(takes) == 1 {
		return takes[0]
	}

	if len(takes) == 2 {
		return takes[0]
	}

	return gametake.PASSE
}

func (p *Player) SetTake(gt *gametake.GameTake) {
	p.Take = gt
}
