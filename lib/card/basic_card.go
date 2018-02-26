package card

import (
	"github.com/c2technology/monopoly-go/lib/game"
)

type BasicCard struct {
	Name   string
	Action func(game.Board, game.Banker, game.Player)
}

func NewCard(name string, action func(game.Board, game.Banker, game.Player)) game.Card {
	return &BasicCard{Name: name, Action: action}
}

func (c *BasicCard) String() string {
	return c.Name

}
func (c *BasicCard) DoAction(board game.Board, banker game.Banker, player game.Player) {
	//TODO: Return the action that the game should run. don't do it here.
	banker.Pay(player, 200)
}
