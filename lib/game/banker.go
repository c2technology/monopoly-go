package game

type Banker interface {
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
