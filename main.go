package main

import (
	"fmt"
	"github.com/c2technology/monopoly-go/lib/card"
	"github.com/c2technology/monopoly-go/lib/dice"
	"github.com/c2technology/monopoly-go/lib/game"
	"github.com/c2technology/monopoly-go/lib/player"
	"log"
	"strconv"
	"github.com/c2technology/monopoly-go/lib/board"
)

func main() {
	log.Printf("Welcome to Monopoly!")
	//TODO: Initialize rentals
	var rentals []game.Rental

	bank := player.NewBanker(1514000, rentals, 32, 12)

	playerCount := getPlayers()

	players := []game.Player{}
	for i := 0; i < playerCount; i++ {
		player := player.NewPlayer(bank)
		bank.Pay(player, 150000)
		players = append(players, player)
	}

	communityChest := card.NewCommunityChest(bank)
	chance := card.NewChance(bank)
	dice := dice.NewDice()
	board := board.NewClassicBoard()
	//TODO: Board initialized

	log.Printf("Initializing game...")
	game := game.NewGame(players, bank, dice, communityChest, chance)

	game.Start()

}

func getPlayers() int {
	log.Printf("Number of player: ")
	text := ""
	fmt.Scanln(&text)
	count, err := strconv.Atoi(text)
	if err != nil || count < 2 {
		log.Printf("Invalid input. Please provide number.")
		return getPlayers()
	}
	return count
}
