package controllers

import (
	"github.com/c2technology/monopoly-go/lib/models"
	"log"
)

type Game struct {
	CommunityChest models.Deck
	Chance models.Deck
	Dice models.Dice
	Players []models.Player
	Bank models.Bank
}

func (g *Game) Start(){
	log.Printf("Let's begin!");
	for len(g.Players) > 0 {
		for i:=0; i < len(g.Players); i++ {
			currentPlayer := g.Players[i]
			log.Printf("Player Turn: %s", currentPlayer);
			g.Dice.Roll()
			log.Printf("%s rolled a %s", currentPlayer, g.Dice)
			if g.Dice.Doubles() {
				log.Printf("%s rolled a %s", currentPlayer, roll)
			}
		}
	}
}