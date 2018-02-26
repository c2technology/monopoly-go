package game

//Card is a single card that can perform an action with the context of a Player
type Card interface {
	//Name of the card
	Name() string

	//Action to perform
	DoAction(GameState, Player)
}

//Card is a single card that can perform an action with the context of a Player
type ConsumableCard interface {
	Card

	//Uses this Card on the given player
	Use(Player)

	//The cash value of this card
	Value() int64
}
