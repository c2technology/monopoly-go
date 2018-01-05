package controllers

import (
	"fmt"
	"log"

	"github.com/c2technology/monopoly-go/lib/models"
)

//Game time!
type Game interface {
	Start()
}

//NewGame creates a new Monopoly game loaded with the given data
func NewGame(players []models.Player, bank models.Bank, dice models.Dice, chance, communityChest models.Deck) Game {
	game := &defaultGame{
		communityChest,
		chance,
		dice,
		players,
		bank,
	}
	fmt.Printf("Dice initialzied: %v", dice)
	return game

}

type defaultGame struct {
	CommunityChest models.Deck
	Chance         models.Deck
	Dice           models.Dice
	Players        []models.Player
	Bank           models.Bank
}

func (g *defaultGame) Start() {
	log.Println("Let's begin!")
	turn := 1
	for len(g.Players) > 0 {
		log.Printf("Begin Turn %d", turn)
		for i := 0; i < len(g.Players); i++ {
			currentPlayer := g.Players[i]
			log.Printf("Player: %s", currentPlayer)
			g.Dice.Roll()
			log.Printf("%s rolled a %s", currentPlayer, g.Dice)
			if g.Dice.Doubles() {
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
