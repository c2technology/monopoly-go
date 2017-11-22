package controllers

import (
	"github.com/c2technology/monopoly-go/lib/models"
	"log"
	"fmt"
)

type Game interface {
	Start()
}

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
	Chance models.Deck
	Dice models.Dice
	Players []models.Player
	Bank models.Bank
}

func (g *defaultGame) Start(){
	log.Println("Let's begin!")
	turn := 1
	for len(g.Players) > 0 {
		log.Printf("Begin Turn %s", turn)
		for i:=0; i < len(g.Players); i++ {
			currentPlayer := g.Players[i]
			log.Printf("Player: %s", currentPlayer)
			g.Dice.Roll()
			log.Printf("%s rolled a %s", currentPlayer, g.Dice)
			if g.Dice.Doubles() {
				log.Printf("%s rolled a double (%s)", currentPlayer, g.Dice)
			}
		}
		log.Printf("End Turn %s", turn)
		turn = turn + 1
	}
}