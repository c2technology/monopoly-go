package models

//Rental is a Space that can be rented
type Rental interface {
	Space
	Group() Group
	Price() int64
	Rent() int64
	Proprietor() Proprietor
	MortgageValue() int64
	Mortgaged() bool
	Mortgage()
	Unmortgage()
	PayoffValue() int64
}

//Buildable is a Rental that can be build upon
type Buildable interface {
	Rental
	Houses() int
	Hotels() int
	BuildHouse()
	BuildHotel()
	DemolishHouse()
	DemolishHotel()
	BuildingPrice() int64
}
