package game

import (
	"fmt"
	"log"
)

//NewGame creates a new Monopoly game loaded with the given data
func NewGame(board Board, players []Player, banker Banker, dice Dice, chance, communityChest Deck) Game {
	game := &defaultGame{
		communityChest,
		chance,
		dice,
		players,
		banker,
		board,
		make(map[Player]int),
	}
	fmt.Printf("Dice initialzied: %v", dice)
	return game

}

type defaultGame struct {
	communityChest Deck
	chance         Deck
	dice           Dice
	players        []Player
	banker         Banker
	board          Board
	doubles        map[Player]int
	//TODO: Game stats? track doubles,
}


func (g *defaultGame) Banker() Banker {
	return g.banker
}

func (g *defaultGame) Players() []Player {
	return g.players
}

func (g *defaultGame) Start() {
	log.Println("Let's begin!")
	turn := 1
	for len(g.players) > 0 {
		log.Printf("Begin Turn %d", turn)
		for i := 0; i < len(g.players); i++ {
			currentPlayer := g.players[i]
			log.Printf("Player: %s", currentPlayer)
			g.dice.Roll()
			log.Printf("%s rolled a %s", currentPlayer, g.dice)
			if g.dice.Doubles() {
				log.Printf("Doubles! %s will go again!", currentPlayer)
				//TODO: Catch 3x doubles in a row and send to jail
			}
			//TODO: Move player g.Dice.FaceValue() spaces
			//TODO: Use player.WantToBuy to determine if player should buy
			//		TODO: Make transaction if true
			//
		}
		log.Printf("End Turn %d", turn)
		turn = turn + 1
	}
}

func (g* defaultGame) CommunityChest() Deck {
	return g.communityChest
}
func (g* defaultGame) Chance() Deck {
	return g.chance
}