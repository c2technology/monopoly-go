package player

import (
	"fmt"
	"github.com/c2technology/monopoly-go/lib/game"
)

//TODO: This could eventually be a real player

type Banker struct {
	Name    string
	Cash    int64
	Rentals []game.Rental
	Houses  int
	Hotels  int
}

func (b *Banker) String() string {
	return fmt.Sprintf("%s: ($%d) - Houses: %d, Hotels: %d, Properties: %d", b.Name, b.Cash, b.Houses, b.Hotels, len(b.Rentals))
}

func (b *Banker) Receive(cash int64, rentals []game.Rental, cards []game.ConsumableCard) {
	b.Cash = b.Cash + cash
	b.Rentals = append(b.Rentals, rentals...)
	for _, card := range cards {
		card.Use(nil)
	}
}

func (b *Banker) Pay(collector game.Collector, cash int64) {
	collector.Receive(cash, nil, nil)
	b.Cash = b.Cash - cash
}

func (b *Banker) SellRental(proprietor game.Proprietor, rental game.Rental) {
	found := false
	for _, owned := range b.Rentals {
		if rental == owned {
			found = true
			break
		}
	}
	if found {
		proprietor.Pay(b, rental.Price())
		proprietor.Receive(0, []game.Rental{rental}, nil)
	}
}

func (b *Banker) BuyHouse(debtor game.Debtor, rental game.Buildable) {
	if b.Houses > 0 {
		debtor.Pay(b, rental.BuildingPrice())
		b.Houses = b.Houses - 1
		rental.BuildHouse()
	}
}

func (b *Banker) BuyHotel(debtor game.Debtor, rental game.Buildable) {
	if b.Hotels > 0 {
		debtor.Pay(b, rental.BuildingPrice())
		b.Hotels = b.Hotels - 1
		rental.BuildHotel()
	}
}

func (b *Banker) SellHouse(collector game.Collector, rental game.Buildable) {
	rental.DemolishHouse()
	collector.Receive(rental.BuildingPrice()/2, nil, nil)
	b.Houses = b.Houses + 1
}

func (b *Banker) SellHotel(collector game.Collector, rental game.Buildable) {
	rental.DemolishHotel()
	collector.Receive(rental.BuildingPrice()/2, nil, nil)
	b.Hotels = b.Hotels + 1
}

func (b *Banker) SellCard(collector game.Collector, card game.ConsumableCard) {
	collector.Receive(card.Value(), nil, nil)
	b.Cash = b.Cash - card.Value()
}

func (b *Banker) Mortgage(collector game.Collector, rental game.Rental) {
	if !rental.Mortgaged() {
		rental.Mortgage()
		collector.Receive(rental.MortgageValue(), nil, nil)
	}
}
func (b *Banker) Unmortgage(debtor game.Debtor, rental game.Rental) {
	if rental.Mortgaged() {
		//TODO: Could lead to bankruptcy to bank if called on accidentally or unnecessarily
		debtor.Pay(b, rental.PayoffValue())
		rental.Mortgage()
	}
}
