package card

import (
	"github.com/c2technology/monopoly-go/lib/game"
)

type basicCard struct {
	name   string
	action func(game.GameState, game.Player)
}

func NewCard(name string, action func(game.GameState, game.Player)) game.Card {
	return &basicCard{name: name, action: action}
}

func (c *basicCard) Name() string {
	return c.name

}
func (c *basicCard) DoAction(gameState game.GameState, player game.Player) {
	gameState.Banker().Pay(player, 200)
}
