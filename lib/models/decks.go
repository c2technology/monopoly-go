package models

//Deck is a collection of Cards
type Deck interface {
	//Shuffle this Deck
	Shuffle()

	//Draws a Card from this Deck and gives it to the Player
	Draw(Player)

}

//Card is a single card that can perform an action with the context of a Player
type Card interface{
	//Name of the card
	Name() string

	//Action to perform
	DoAction(Player)
}

//Card is a single card that can perform an action with the context of a Player
type ConsumableCard interface{
	Card

	//Uses this Card on the given player
	Use(Player)

	//The cash value of this card
	Value() int64
}


