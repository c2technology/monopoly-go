package game

import "fmt"

//Card is a single card that can perform an action with the context of a Player
type Card interface {
	fmt.Stringer

	//Action to perform
	DoAction(Board, Banker, Player)
}

//Card is a single card that can perform an action with the context of a Player
type ConsumableCard interface {
	Card

	//Uses this Card on the given player
	Use(Player)

	//The cash value of this card
	Value() int64
}
