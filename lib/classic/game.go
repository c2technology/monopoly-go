package classic

import (
	"fmt"
	"log"
	"github.com/c2technology/monopoly-go/lib/game"
)

//Game creates a new Monopoly game loaded with the given data
func Game(board game.Board, players []game.Player, banker game.Banker, dice game.Dice, chance, communityChest game.Deck) game.Game {
	game := &defaultGame{
		communityChest,
		chance,
		dice,
		players,
		banker,
		board,
		make(map[game.Player]int),
	}
	fmt.Printf("Dice initialzied: %v", dice)
	return game

}

type defaultGame struct {
	communityChest game.Deck
	chance         game.Deck
	dice           game.Dice
	players        []game.Player
	banker         game.Banker
	board          game.Board
	doubles        map[game.Player]int
	//TODO: Game stats? track doubles,
}


func (g *defaultGame) Banker() game.Banker {
	return g.banker
}

func (g *defaultGame) Players() []game.Player {
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

func (g* defaultGame) CommunityChest() game.Deck {
	return g.communityChest
}
func (g* defaultGame) Chance() game.Deck {
	return g.chance
}