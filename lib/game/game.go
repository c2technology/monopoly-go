package game

//Game time!
type Game interface {
	Start()
	Banker() Banker
	Players() []Player
	CommunityChest() Deck
	Chance() Deck
}
