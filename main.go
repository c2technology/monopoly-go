package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/c2technology/monopoly-go/lib/controllers"
	"github.com/c2technology/monopoly-go/lib/models"
)

func main() {
	log.Printf("Welcome to Monopoly!")
	//TODO: Initialiez rentals
	rentals := []models.Rental{}

	bank := models.NewBank(1514000, rentals, 32, 12)

	playerCount := getPlayers()

	players := []models.Player{}
	for i := 0; i < playerCount; i++ {
		player := models.NewPlayer(bank)
		bank.Pay(player, 150000)
		players = append(players, player)
	}

	communityChest := models.NewCommunityChest()
	chance := models.NewChance()
	dice := models.NewDice()

	//TODO: Board initialized

	log.Printf("Initializing game...")
	game := controllers.NewGame(players, bank, dice, communityChest, chance)

	game.Start()

}

func getPlayers() int {
	log.Printf("Number of players: ")
	text := ""
	fmt.Scanln(&text)
	count, err := strconv.Atoi(text)
	if err != nil || count < 2 {
		log.Printf("Invalid input. Please provide number.")
		return getPlayers()
	}
	return count
}
