package board

import (
	"encoding/json"
	"fmt"
	"github.com/c2technology/monopoly-go/lib/game"
)

func NewActionableSpace(name string, action func(game.Player)) game.Space {
	return &actionSpace{
		name:          name,
		action: action,
	}
}

type actionSpace struct {
	name          string
	action func (game.Player)
}

func (r *actionSpace) String() string {
	return fmt.Sprintf("%v", r.name)
}
func (r *actionSpace) DoAction(player game.Player) {
	r.action(player)
}

