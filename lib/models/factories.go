package models

import "math/rand"

func NewCommunityChest() Deck {
	//TODO: Make cards
	return &basicDeck{name: "Community Chest"}
}

func NewChance() Deck {
	//TODO: Make cards
	return &basicDeck{name: "Chance"}
}

	//Creates a new Player
func NewPlayer(bank Bank) Player {
	return &defaultPlayer{
		name: playerNames[rand.Intn(len(playerNames))],
		bank: bank,
		//TODO: Get random strategy
	}
}

func NewDice() Dice {
	return &dice{
		die1: newDie(6),
		die2: newDie(6),
	}

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
	return &standardBank{"Standard Bank", cash, rentals, houses,  hotels}
}
