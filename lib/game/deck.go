package game

//Deck is a collection of Cards
type Deck interface {
	//Shuffle this Deck
	Shuffle()

	//Draws a Card from this Deck and gives it to the Player
	Draw(Player)
}
