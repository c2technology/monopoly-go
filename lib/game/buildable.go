package game

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
