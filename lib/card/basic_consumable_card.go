package card

import (
	"github.com/c2technology/monopoly-go/lib/game"
)

type basicConsumableCard struct {
	*basicCard
	usage func(game.GameState, game.Player)
	value int64
}

func NewConsumableCard(name string, action func(game.GameState, game.Player), usage func(game.Player), value int64) game.Card {
	card := &basicCard{name: name, action: action}
	return &basicConsumableCard{basicCard: card, usage: usage, value: value}
}

func (c *basicConsumableCard) Value() int64 {
	return c.value

}
func (c *basicConsumableCard) Use(gameState game.GameState, player game.Player) {
	c.usage(gameState, player)
}
