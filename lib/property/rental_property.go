package property

import (
	"fmt"
	"github.com/c2technology/monopoly-go/lib/game"
)

type Rent struct {
	rent map[int]int64
}

func NewRental(name string, group game.Group, banker game.Banker, price, mortgage, buildingPrice int64, rent map[int]int64) game.Rental {
	return &rentalProperty{
		name:          name,
		group:         group,
		price:         price,
		rent:          rent,
		owner:         banker,
		mortgage:      mortgage,
		buildingPrice: buildingPrice,
	}
}

type rentalProperty struct {
	name          string
	group         game.Group
	price         int64
	rent          map[int]int64
	owner         game.Proprietor
	houses        int
	hotels        int
	mortgage      int64
	mortgaged     bool
	buildingPrice int64
}

func (r *rentalProperty) String() string {
	return fmt.Sprintf("%v %v - %v", r.name, r.group, r.owner)
}
func (r *rentalProperty) DoAction(player game.Player) {
	if bank, ok := r.Proprietor().(game.Banker); ok {
		//bank owns this
		if player.WantToBuy(r) {
			bank.SellRental(player, r)
			//} else {
			//TODO: Auction
		}
	} else {
		player.Pay(r.owner, r.rent[r.houses+r.hotels])
	}
}
func (r *rentalProperty) Group() game.Group {
	return r.group
}
func (r *rentalProperty) BuildingPrice() int64 {
	return r.buildingPrice
}
func (r *rentalProperty) Price() int64 {
	return r.price
}
func (r *rentalProperty) Rent() int64 {
	return r.rent[r.houses+r.hotels]
}
func (r *rentalProperty) Proprietor() game.Proprietor {
	return r.owner
}
func (r *rentalProperty) BuildHouse() {
	r.houses = r.houses + 1
}
func (r *rentalProperty) BuildHotel() {
	r.hotels = r.hotels + 1
}
func (r *rentalProperty) DemolishHouse() {
	r.houses = r.houses - 1
}
func (r *rentalProperty) DemolishHotel() {
	r.hotels = r.hotels - 1
}

func (r *rentalProperty) Houses() int {
	return r.houses
}
func (r *rentalProperty) Hotels() int {
	return r.hotels
}
func (r *rentalProperty) Mortgage() {
	r.mortgaged = true
}
func (r *rentalProperty) Unmortgage() {
	r.mortgaged = false
}
func (r *rentalProperty) MortgageValue() int64 {
	return r.mortgage
}
func (r *rentalProperty) Mortgaged() bool {
	return r.mortgaged
}

func (r *rentalProperty) PayoffValue() int64 {
	return int64(float64(r.mortgage) * 1.2)
}
