package game

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"pathe.co/zinx/gametake"
	"pathe.co/zinx/pkg/cards"
	"pathe.co/zinx/pkg/player"
)

var (
	ErrCardsAlreadyDispatched   = errors.New("cards already dispatched error")
	ErrCardNotFoundInPlayerHand = errors.New("card not found in player hand")
	ErrDeckNotFound             = errors.New("deck not found")
	ErrGameIsFull               = errors.New("Game is full")
	ErrDuplicatePlayerTake      = errors.New("oops duplicate player take")
	ErrBadTake                  = errors.New("oops bad take, choose a greater take or pass")
)

type Game struct {
	CartesDistribuees      int
	nombrePli              int
	pliCardsCount          int
	NombreJoueurs          int
	TakesFinished          bool
	scoreTeamA, scoreTeamB int
	Cartes                 [32]cards.Card
	Decks                  [8]Deck
	Plis                   [8][4]cards.Card
	players                [4]*player.Player
	ring                   [4]int
	take                   gametake.GameTake
}

func NewGame() *Game {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	jeu := cards.CardSet{Cards: [32]cards.Card{}}

	newGame := Game{
		Cartes:            jeu.Distribute(),
		Plis:              [8][4]cards.Card{},
		Decks:             [8]Deck{},
		CartesDistribuees: 0,
		NombreJoueurs:     0,
		TakesFinished:     false,
		players:           [4]*player.Player{},
		take:              gametake.PASSE,
		nombrePli:         0,
		pliCardsCount:     0,
		ring:              [4]int{0, 1, 2, 3},
	}

	return &newGame
}

func (g *Game) Score() (int, int) {
	return g.scoreTeamA, g.scoreTeamB
}

func (g *Game) NextRound(playerID int) [4]int {
	switch playerID {
	case 0:
		return [4]int{0, 1, 2, 3}
	case 1:
		return [4]int{1, 2, 3, 0}
	case 2:
		return [4]int{2, 3, 0, 1}
	case 3:
		return [4]int{3, 0, 1, 2}
	default:
		panic("invalid input")
	}
}

func (g *Game) PlayCardNext(playerID int, c cards.Card) error {
	plyr := g.players[playerID]
	hasCard, idx := plyr.HasCard(c)

	if !hasCard {
		return ErrCardNotFoundInPlayerHand
	}

	fmt.Println("Game ring debug before processing", g.ring)
	if g.Decks[g.nombrePli].cardscount == 0 {
		g.Decks[g.nombrePli] = NewDeck(g.ring, g.take)
	}

	if err := g.Decks[g.nombrePli].AddCard(playerID, c); err != nil {
		return err
	}

	plyr.PlayingHand.Cards[idx] = cards.Card{Genre: "", Couleur: ""}
	if g.Decks[g.nombrePli].cardscount == 4 {
		fmt.Println("Winner debug", g.Decks[g.nombrePli].winner)
		fmt.Println("Winner debug", g.Decks[g.nombrePli])
		g.ring = g.NextRound(g.Decks[g.nombrePli].winner)
		fmt.Println("Game ring debug", g.ring)
		a, b := g.Decks[g.nombrePli].Score()
		g.scoreTeamA += a
		g.scoreTeamB += b
		g.nombrePli++
	}

	fmt.Println(g.ring)

	return nil
}

func (g *Game) PlayCard(playerID int, card cards.Card) error {
	plyr := g.players[playerID]
	hasCard, idx := plyr.HasCard(card)

	if !hasCard {
		return ErrCardNotFoundInPlayerHand
	}

	g.Plis[g.nombrePli][g.pliCardsCount] = card
	plyr.PlayingHand.Cards[idx] = cards.Card{Genre: "", Couleur: ""}

	g.pliCardsCount++

	if g.pliCardsCount == 4 {
		g.nombrePli++
		g.pliCardsCount = 0
	}

	return nil
}

func (g *Game) CurrentDeck() ([4]cards.Card, error) {
	if g.nombrePli > 7 {
		return [4]cards.Card{}, ErrDeckNotFound
	}

	return g.Decks[g.nombrePli].cards, nil
}

func (g *Game) AddPlayer(plyr *player.Player) error {
	if g.NombreJoueurs == 4 {
		return ErrGameIsFull
	}

	plyr.Hand.Cards = g.distribute()
	plyr.SetID(g.NombreJoueurs)
	g.players[g.NombreJoueurs] = plyr
	g.NombreJoueurs++

	return nil
}

func (g *Game) AddTake(playerID int, take gametake.GameTake) error {
	if g.players[playerID].Take != nil {
		return ErrDuplicatePlayerTake
	}

	g.players[playerID].Take = &take

	if g.take.GreaterThan(take) && take != gametake.PASSE {
		return ErrBadTake
	}

	if g.take == gametake.PASSE {
		g.take = take
	} else if take != gametake.PASSE {
		g.take = take
	}

	if g.take == gametake.TOUT || g.takesComplete() {
		g.TakesFinished = true

		if err := g.DispatchCards(); err != nil {
			return err
		}

		g.sendPlayingHands()
	}

	return nil
}

func (g *Game) sendPlayingHands() {
	for _, plyr := range g.players {
		if plyr != nil {
			fmt.Println("sending playing hand to player")

			r := ReceivePlayingHandEvt(*plyr, g.GetTake())
			jsonMsg, err := json.Marshal(r)

			if err != nil {
				fmt.Println(err)
			}

			if plyr.Conn != nil {
				if err := plyr.Conn.WriteMessage(1, jsonMsg); err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}

func (g *Game) takesComplete() bool {
	for _, p := range g.players {
		if p != nil && p.Take == nil {
			return false
		}
	}

	return true
}

func (g *Game) distribute() [5]cards.Card {
	cards := [5]cards.Card{}
	for i := 0; i < 5; i++ {
		cards[i] = g.Cartes[g.CartesDistribuees+i]
	}

	g.CartesDistribuees += 5

	return cards
}

func (g *Game) GetPlayers() [4]*player.Player {
	return g.players
}

func (g *Game) GetTake() gametake.GameTake {
	return g.take
}

func (g *Game) AvailableTakes() []gametake.GameTake {
	takes := []gametake.GameTake{}
	takes = append(takes, gametake.PASSE)

	for _, t := range gametake.AllTakes {
		if t.GreaterThan(g.take) {
			takes = append(takes, t)
		}
	}

	return takes
}

func (g *Game) DispatchCards() error {
	if g.CartesDistribuees == 32 {
		return ErrCardsAlreadyDispatched
	}

	for _, plyr := range g.players {
		cards := []cards.Card{}

		if plyr != nil {
			for _, c := range plyr.Hand.Cards {
				cards = append(cards, c)
			}

			for i := 0; i < 3; i++ {
				cards = append(cards, g.Cartes[g.CartesDistribuees])
				g.CartesDistribuees++
			}

			plyr.PlayingHand = player.PlayingHand{Cards: cards}
		}
	}

	return nil
}
