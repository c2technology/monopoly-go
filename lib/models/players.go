package models

import (
	"fmt"
)

//Someone who owes something
type Debtor interface {
	//Pays the given amount to the given Collector. Steps to ensure payment is made should be performed.
	Pay(Collector, int64)
}

//Collects assets
type Collector interface {
	//Receive assets as a one sided-exchange
	Receive(int64, []Rental, []ConsumableCard)
}

type Proprietor interface {
	Debtor
	Collector
}

type Buyer interface {
	WantToBuy(Rental) bool
}

type Prisoner interface {
	InJail() bool
	SentanceRemaining() int
	PostBail()
}

type Builder interface {
	//Builds a building in the given Rental Group.
	Build(Group)
}

type Player interface {
	fmt.Stringer
	Proprietor
	Builder
	Prisoner
	Buyer
}

type Bank interface {
	Proprietor
	SellRental(Proprietor, Rental)
	BuyHouse(Debtor, Buildable)
	BuyHotel(Debtor, Buildable)
	SellHouse(Collector, Buildable)
	SellHotel(Collector, Buildable)
	SellCard(Collector, ConsumableCard)
	Mortgage(Collector, Rental)
	Unmortgage(Debtor, Rental)
}
