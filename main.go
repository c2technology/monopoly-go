package main

import (
	"fmt"
	"github.com/c2technology/monopoly-go/lib/game"
	"github.com/c2technology/monopoly-go/lib/player"
	"log"
	"strconv"
	"github.com/c2technology/monopoly-go/lib/classic"
)

func main() {
	log.Printf("Welcome to Monopoly!")
	//TODO: Initialize rentals
	var rentals []game.Rental

	bank := &player.Banker{Name: "Banker", Cash: 1514000, Rentals: rentals, Houses: 32, Hotels: 12}

	playerCount := getPlayers()

	var players []game.Player
	for i := 0; i < playerCount; i++ {
		p := player.RandomPlayer(bank)
		bank.Pay(p, 150000)
		players = append(players, p)
	}

	//TODO: Factory with different styles of Monopoly (classic, kid, holiday, etc)
	board := classic.Board()
	communityChest := classic.CommunityChest(board, bank)
	chance := classic.Chance(board, bank)
	dice := classic.Dice()

	//TODO: board initialized

	log.Printf("Initializing game...")
	game := classic.Game(board, players, bank, dice, communityChest, chance)
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
