package game

import (
	"fmt"
)

//Space is a space on the Board. It could be a Property or could be an Action
type Space interface {
	fmt.Stringer
	DoAction(Player)
}
