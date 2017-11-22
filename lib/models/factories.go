package models

import (
	"math/rand"
	"fmt"
)

func NewCommunityChest() Deck {
	//TODO: Make cards
	deck := &basicDeck{name: "Community Chest"}
	fmt.Printf("Initializing %s", deck)
	return deck
}

func NewChance() Deck {
	//TODO: Make cards
	deck := &basicDeck{name: "Chance"}
	fmt.Printf("Initializing %s", deck)
	return deck
}

	//Creates a new Player
func NewPlayer(bank Bank) Player {
	player := &defaultPlayer{
		name: playerNames[rand.Intn(len(playerNames))],
		bank: bank,
		//TODO: Get random strategy
	}
	fmt.Printf("%s has taken a seat!", player)
	return player
}



func NewDice() Dice {
	dice := &dice{
		die1: newDie(6),
		die2: newDie(6),
	}
	fmt.Printf("Dice initialzied: %v", dice)
	return dice

}
func newDie(values int) Die {
	if values < 3 {
		values = 3
	}
	vals := []int{}
	for i := 1; i <= values; i++ {
		vals[i-1] = i
	}
	current := vals[rand.Intn(len(vals))]
	return &die{
		vals,
		current,
	}
}

func NewBank(cash int64, rentals []Rental, houses, hotels int) Bank {
	//TODO: Load all the rentals
	fmt.Println("Initializing Bank...")
	return &standardBank{"Standard Bank", cash, rentals, houses,  hotels}
}

//func NewRental(name string, group Group, price, rent,  mortgage int64) Rental {
//	return &rentalProperty{
//		name:     name,
//		group:    group,
//		price:    price,
//		rent:     rent,
//		mortgage: mortgage,
//	}
//}
