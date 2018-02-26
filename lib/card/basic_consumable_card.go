package card

import (
	"github.com/c2technology/monopoly-go/lib/game"
)

type BasicConsumableCard struct {
	*BasicCard
	Usage func(game.Player)
	SaleValue int64
}


func (c *BasicConsumableCard) Value() int64 {
	return c.SaleValue
}
func (c *BasicConsumableCard) Use(player game.Player) {
	c.Usage(player)
}
//TODO: Sell, Trade?
