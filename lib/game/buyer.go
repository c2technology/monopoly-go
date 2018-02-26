package game

type Buyer interface {
	WantToBuy(Rental) bool
}
