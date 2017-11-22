package main

import (
	"fmt"
	"github.com/c2technology/monopoly-go/lib/models"
	"os"
	"bufio"
	"strconv"
	"github.com/c2technology/monopoly-go/lib/controllers"
)

func main() {
	fmt.Println("Welcome to Monopoly!")
	//TODO: Initialiez rentals
	rentals := []models.Rental{}

	bank := models.NewBank(1514000, rentals, 32,  12)

	playerCount := getPlayers()

	players := []models.Player{}
	for i := 0;	i < playerCount; i++	{
		player := models.NewPlayer(bank)
		bank.Pay(player, 150000)
		players[i] = player
	}

	communityChest := models.NewCommunityChest()
	chance := models.NewChance()
	dice := models.NewDice()

	//TODO: Board initialized

	fmt.Println("Initializing game...")
	game := controllers.NewGame(players, bank, dice, communityChest, chance)

	game.Start()

}

func getPlayers() int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Number of players: ")
	text, _ := reader.ReadString('\n')
	if count, err := strconv.Atoi(text); err != nil {
		fmt.Println("Invalid input. Please provide number.")
		return getPlayers()
	} else {
		return count
	}
}