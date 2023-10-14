package cards

import (
	"fmt"
	"math/rand"
	"time"
)

var Jeu32Cartes [32]Card = [32]Card{
	{"A", "Pique"},
	{"A", "Coeur"},
	{"A", "Carreau"},
	{"A", "Trefle"},
	{"V", "Pique"},
	{"V", "Coeur"},
	{"V", "Carreau"},
	{"V", "Trefle"},
	{"R", "Pique"},
	{"R", "Coeur"},
	{"R", "Carreau"},
	{"R", "Trefle"},
	{"D", "Pique"},
	{"D", "Coeur"},
	{"D", "Carreau"},
	{"D", "Trefle"},
	{"10", "Pique"},
	{"10", "Coeur"},
	{"10", "Carreau"},
	{"10", "Trefle"},
	{"9", "Pique"},
	{"9", "Coeur"},
	{"9", "Carreau"},
	{"9", "Trefle"},
	{"8", "Pique"},
	{"8", "Coeur"},
	{"8", "Carreau"},
	{"8", "Trefle"},
	{"7", "Pique"},
	{"7", "Coeur"},
	{"7", "Carreau"},
	{"7", "Trefle"},
}

type Card struct {
	Genre   string
	Couleur string
}

func (c *Card) String() string {
	return fmt.Sprintf("(%s %s)", c.Genre, c.Couleur)
}

func (c *Card) IsValet() bool {
	return c.Genre == "V"
}

func (c *Card) IsAce() bool {
	return c.Genre == "A"
}

func (c *Card) IsNine() bool {
	return c.Genre == "9"
}

type CardSet struct {
	Cards [32]Card
}

func (j CardSet) Distribuer() [32]Card {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	cartes := CardSet{Jeu32Cartes}
	rand.Shuffle(len(cartes.Cards), func(i, j int) {
		cartes.Cards[i], cartes.Cards[j] = cartes.Cards[j], cartes.Cards[i]
	})

	return cartes.Cards
}
