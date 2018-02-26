package game

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
