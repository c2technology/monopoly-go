package models

import (
	"fmt"
	"log"
	"math/rand"
)

func NewCommunityChest() Deck {
	//TODO: Make cards
	deck := &basicDeck{name: "Community Chest"}
	log.Printf("Initializing %s", deck)
	return deck
}

func NewChance() Deck {
	//TODO: Make cards
	deck := &basicDeck{name: "Chance"}
	log.Printf("Initializing %s", deck)
	return deck
}

//Creates a new Player
func NewPlayer(bank Bank) Player {
	player := &defaultPlayer{
		name: playerNames[rand.Intn(len(playerNames))],
		bank: bank,
		//TODO: Get random strategy
	}
	log.Printf("%s has taken a seat!", player)
	return player
}

func NewDice() Dice {
	dice := &dice{
		die1: newDie(6),
		die2: newDie(6),
	}
	log.Printf("Dice initialzied: %v", dice)
	return dice

}
func newDie(sides int) Die {
	if sides < 3 {
		sides = 3
	}
	vals := []int{}
	for i := 0; i < sides; i++ {
		vals = append(vals, i+1)
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
	return &standardBank{"Standard Bank", cash, rentals, houses, hotels}
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
