package card

import (
	"fmt"
	"github.com/c2technology/monopoly-go/lib/game"
	"log"
	"math/rand"
)

type BasicDeck struct {
	Name  string
	Cards []game.Card
	Board game.Board
	Banker game.Banker
}

func (c *BasicDeck) String() string {
	return fmt.Sprintf("%s: %v", c.Name, len(c.Cards))
}

func (c *BasicDeck) Shuffle() {
	log.Printf("%s was shuffled!", c)
	deck := make([]game.Card, len(c.Cards))
	perm := rand.Perm(len(c.Cards))
	for i, v := range perm {
		deck[v] = c.Cards[i]
	}
	c.Cards = deck

}

func (c *BasicDeck) Draw(player game.Player) {
	log.Printf("%s draws a %s card", player, c)
	card := c.Cards[0]
	if consumableCard, ok := card.(game.ConsumableCard); ok {
		player.Receive(0, nil, []game.ConsumableCard{consumableCard})
		return
	}
	card.DoAction(c.Board, c.Banker, player)
}
