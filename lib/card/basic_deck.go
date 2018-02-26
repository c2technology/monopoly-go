package card

import (
	"fmt"
	"github.com/c2technology/monopoly-go/lib/game"
	"log"
	"math/rand"
)

type basicDeck struct {
	name  string
	cards []game.Card
}

func NewBasicDeck(name  string, cards []game.Card) game.Deck {
	deck := &basicDeck{name: "Chance", cards: cards}
	deck.Shuffle()
	return deck
}
func (c *basicDeck) String() string {
	return fmt.Sprintf("%s: %v", c.name, len(c.cards))
}

func (c *basicDeck) Shuffle() {
	log.Printf("%s was shuffled!", c)
	dest := make([]game.Card, len(c.cards))
	perm := rand.Perm(len(c.cards))
	for i, v := range perm {
		dest[v] = c.cards[i]
	}
	c.cards = dest

}

func (c *basicDeck) Draw(player game.Player) {
	log.Printf("%s draws a %s card", player, c)
	card := c.cards[0]
	if consumableCard, ok := card.(game.ConsumableCard); ok {
		player.Receive(0, nil, []game.ConsumableCard{consumableCard})
		return
	}
	card.DoAction(gameState, player)
}
