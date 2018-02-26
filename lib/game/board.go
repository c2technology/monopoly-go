package game

type Board interface {

	//Gets the Space the given player is currently in
	Space(Player) Space

	//Moves the player to the given location on the Board and returns the Space landed upon
	Move(Player, int) Space
}
