package models

import "fmt"

type Group string

const (
	GREEN      = Group("GREEN")
	PURPLE     = Group("PURPLE")
	ORANGE     = Group("ORANGE")
	RED        = Group("RED")
	YELLOW     = Group("YELLOW")
	BLUE       = Group("BLUE")
	LIGHT_BLUE = Group("LIGHT_BLUE")
	BROWN      = Group("BROWN")
	RAILROAD   = Group("RAILROAD")
	UTILITY    = Group("UTILITY")
)

//Space is a space on the Board. It could be a Property or could be an Action
type Space interface {
	fmt.Stringer
	DoAction(Player) error
}

type Board interface {
}
