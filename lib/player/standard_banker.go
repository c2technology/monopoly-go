package player

import (
	"fmt"
	"github.com/c2technology/monopoly-go/lib/game"
)

func NewBanker(cash int64, rentals []game.Rental, houses, hotels int) game.Banker {
	//TODO: Load all the rentals
	fmt.Println("Initializing Bank...")
	return &standardBanker{"Standard Bank", cash, rentals, houses, hotels}
}

type standardBanker struct {
	name    string
	cash    int64
	rentals []game.Rental
	houses  int
	hotels  int
}

func (b *standardBanker) String() string {
	return fmt.Sprintf("%s: ($%d) - Houses: %d, Hotels: %d, Properties: %d", b.name, b.cash, b.houses, b.hotels, len(b.rentals))
}

func (b *standardBanker) Receive(cash int64, rentals []game.Rental, cards []game.ConsumableCard) {
	b.cash = b.cash + cash
	b.rentals = append(b.rentals, rentals...)
	for _, card := range cards {
		card.Use(nil)
	}
}

func (b *standardBanker) Pay(collector game.Collector, cash int64) {
	collector.Receive(cash, nil, nil)
	b.cash = b.cash - cash
}

func (b *standardBanker) SellRental(proprietor game.Proprietor, rental game.Rental) {
	found := false
	for _, owned := range b.rentals {
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

func (b *standardBanker) BuyHouse(debtor game.Debtor, rental game.Buildable) {
	if b.houses > 0 {
		debtor.Pay(b, rental.BuildingPrice())
		b.houses = b.houses - 1
		rental.BuildHouse()
	}
}

func (b *standardBanker) BuyHotel(debtor game.Debtor, rental game.Buildable) {
	if b.hotels > 0 {
		debtor.Pay(b, rental.BuildingPrice())
		b.hotels = b.hotels - 1
		rental.BuildHotel()
	}
}

func (b *standardBanker) SellHouse(collector game.Collector, rental game.Buildable) {
	rental.DemolishHouse()
	collector.Receive(rental.BuildingPrice()/2, nil, nil)
	b.houses = b.houses + 1
}

func (b *standardBanker) SellHotel(collector game.Collector, rental game.Buildable) {
	rental.DemolishHotel()
	collector.Receive(rental.BuildingPrice()/2, nil, nil)
	b.hotels = b.hotels + 1
}

func (b *standardBanker) SellCard(collector game.Collector, card game.ConsumableCard) {
	collector.Receive(card.Value(), nil, nil)
	b.cash = b.cash - card.Value()
}

func (b *standardBanker) Mortgage(collector game.Collector, rental game.Rental) {
	if !rental.Mortgaged() {
		rental.Mortgage()
		collector.Receive(rental.MortgageValue(), nil, nil)
	}
}
func (b *standardBanker) Unmortgage(debtor game.Debtor, rental game.Rental) {
	if rental.Mortgaged() {
		//TODO: Could lead to bankruptcy to bank if called on accidentally or unnecessarily
		debtor.Pay(b, rental.PayoffValue())
		rental.Mortgage()
	}
}
