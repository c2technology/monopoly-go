package board

import (
	"fmt"
	"github.com/c2technology/monopoly-go/lib/game"
)

type ActionSpace struct {
	Name          string
	Action func (game.Player)
}

func (r *ActionSpace) String() string {
	return fmt.Sprintf("%v", r.Name)
}
func (r *ActionSpace) DoAction(player game.Player) {
	r.Action(player)
}

