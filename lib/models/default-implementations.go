package models

import "fmt"
import (
	"math/rand"
	"log"
)

type rentalProperty struct {
	name      string
	group     Group
	price     int64
	rent      int64
	owner     Proprietor
	houses    int
	hotels    int
	mortgage  int64
	mortgaged bool
	buildingPrice int64
}
func (r *rentalProperty) String() string {
	return fmt.Sprintf("%v %v - %v", r.name, r.group, r.owner)
}
func (r *rentalProperty) DoAction(player Player) error {
	if bank, ok := r.Proprietor().(Bank); ok {
		//bank owns this
		if player.WantToBuy(r) {
			bank.SellRental(player, r)
			return nil
		} else {
			//TODO: Auction
		}
	} else {
		player.Pay(r.owner, r.rent)
	}
	return nil
}
func (r *rentalProperty) Group() Group {
	return r.group
}
func (r *rentalProperty) BuildingPrice() int64 {
	return r.buildingPrice
}
func (r *rentalProperty) Price() int64 {
	return r.price
}
func (r *rentalProperty) Rent() int64 {
	return r.rent
}
func (r *rentalProperty) Proprietor() Proprietor {
	return r.owner
}
func (r *rentalProperty) BuildHouse() {
	r.houses = r.houses+1
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
	return r.mortgage * 1.2
}

type basicDeck struct {
	name string
	cards []Card
}

func (c * basicDeck) Name() string {
	return c.name
}

func (c *basicDeck) Shuffle(){
	log.Printf("%s was shuffled!", c)
	dest := make([]Card, len(c.cards))
	perm := rand.Perm(len(c.cards))
	for i, v := range perm {
		dest[v] = c.cards[i]
	}
	c.cards = dest

}

func (c *basicDeck) Draw(player Player) {
	log.Printf("%s draws a %s card", player, c)
	card := c.cards[0]
	if consumableCard, ok := card.(ConsumableCard); ok {
		player.Receive(0, nil, []ConsumableCard{consumableCard})
		return
	}
	card.DoAction(player)
}

type dice struct {
	die1 Die
	die2 Die
	doubles bool
	faceValue int
}

func (d *dice) Roll() {
	d.die1.Roll()
	d.die2.Roll()
	d.doubles = d.die1.Value() == d.die2.Value()
	d.faceValue = d.die1.Value() + d.die2.Value()
}
func (d *dice) Value() int {
	return d.faceValue
}

func (d *dice) Doubles() bool {
	return d.doubles
}
func (d *dice) String() string {
	return fmt.Sprintf("%s + %s = %s (%s)", d.die1, d.die2, d.faceValue, d.doubles)
}

type die struct {
	values []int
	faceValue int
}

func (d *die) String() string {
	return fmt.Sprintf("%s", d.faceValue)
}

func (d *die) Roll() {
	d.faceValue = d.values[rand.Intn(len(d.values))]
}

func (d *die) Value() int {
	return d.faceValue
}

//Player is a member of the game
type defaultPlayer struct {
	bank Bank
	cash int64
	name string
	rentals []Rental
	cards []ConsumableCard
	//	piece GamePiece
	inJail bool
	jailRemaining int
	//	gm GameMaster
}

func (p *defaultPlayer) Pay(collector Collector, owed int64) {
	//figure out if we have enough cash
	if owed > p.cash {
		//start mortgaging properties
		//TODO Implement variable liquidation strategy pattern
		p.liquidationStrategy(owed)

		if owed > p.cash {
			//Not enough liquidity!
			p.declareBankruptcy(collector)
			return
		}
	}
	//we scrounged it all up!
	collector.Receive(owed, nil, nil)
	p.cash = p.cash - owed
}

func (p *defaultPlayer) liquidationStrategy(owed int64) {
	if owed > p.cash {
		//liquidate cards
		for _, card := range p.cards {
			p.bank.SellCard(p, card)
			remove(p.cards, card)
		}
		//mortgage properties
		groups := make(map[Group]interface{})
		for _, rental := range p.rentals {
			groups[rental.Group()] = nil
			p.bank.Mortgage(p, rental)
			if owed <= p.cash {
				return
			}
		}

		//liquidate buildings one at a time until we have enough cash
		for group := range groups {
			for p.liquidateBuilding(group) {
				if owed <= p.cash {
					return
				}
			}
		}
		//run again to catch new properties available for mortgage
		for _, rental := range p.rentals {
			p.bank.Mortgage(p, rental)
			if owed <= p.cash {
				return
			}
		}
	}
}


func (p *defaultPlayer) Receive(cash int64, rentals []Rental, cards []ConsumableCard){
	if cash > 0 {
		p.cash = p.cash + cash
	}
	if rentals != nil {
		p.rentals = append(p.rentals, rentals...)
	}
	if cards != nil {
		p.cards = append(p.cards, cards...)
	}
}

func (p *defaultPlayer) WantToBuy(rental Rental) bool {
	//TODO: Implement BuyerStrategy
	return true
}

func (p *defaultPlayer) InJail() bool{
	return p.inJail
}

func (p *defaultPlayer) SentanceRemaining() int{
	return p.jailRemaining
}
func (p *defaultPlayer)PostBail(){
	//TODO

}

func (p *defaultPlayer) Build(group Group) {
	//TODO: Build evenly

}

func (p *defaultPlayer) liquidateBuilding(group Group) bool {
	var groupRentals []Rental

	//Get the full group
	for _, rental := range p.rentals {
		if rental.Mortgaged() {
			continue
		}
		if rental.Group() == group {
			groupRentals = append(groupRentals, rental)
		}
	}
	if len(groupRentals) < 1 {
		//nothing to sell
		return false
	}

	sortedRentals := []Rental{}
	maxBuildings := int(0)
	for _, rental := range groupRentals {
		buildable, ok := rental.(Buildable)
		if ok {
			buildings := buildable.Houses() + buildable.Hotels()
			if buildings > maxBuildings {
				maxBuildings = buildings
				sortedRentals[maxBuildings]=rental
			}
		}
	}

	//Sell a building from a rental with the most amount of buildings
	buildable, ok := sortedRentals[maxBuildings].(Buildable)
	if ok {
		if buildable.Hotels() >0 {
			p.bank.SellHotel(p, buildable)
		} else {
			p.bank.SellHouse(p, buildable)
		}
	}
	return true
}

func (p *defaultPlayer) declareBankruptcy(collector Collector){
	collector.Receive(p.cash, p.rentals, p.cards)
	p.cash =0
	p.rentals = nil
	p.cards = nil
	//TODO: gm.Eject(p) <-- gm should have the only reference to this player

}

func (p *defaultPlayer) String() string {
	return fmt.Sprintf("%s ($%v)", p.name, p.cash/100)
}


type standardBank struct {
	name string
	cash int64
	rentals []Rental
	houses int
	hotels int
}

func (b * standardBank) String() string{
	return fmt.Sprintf("%s: ($%s) - Houses: %s, Hotels: %s, Properties: %s", b.name, b.cash, b.houses, b.hotels, len(b.rentals))
}

func (b *standardBank) Receive(cash int64, rentals []Rental, cards []ConsumableCard) {
	b.cash = b.cash + cash
	b.rentals = append(b.rentals, rentals...)
	for _, card := range cards {
		card.Use(nil)
	}
}

func (b *standardBank) Pay(collector Collector, cash int64) {
	collector.Receive(cash, nil, nil)
	b.cash = b.cash - cash
}

func (b *standardBank) SellRental(proprietor Proprietor, rental Rental){
	found := false
	for _, owned := range b.rentals {
		if rental == owned {
			found = true
			break
		}
	}
	if found {
		proprietor.Pay(b, rental.Price())
		proprietor.Receive(0, []Rental{rental}, nil)
	}
}

func (b *standardBank) BuyHouse(debtor Debtor,rental Buildable){
	if b.houses > 0 {
		debtor.Pay(b, rental.BuildingPrice())
		b.houses = b.houses - 1
		rental.BuildHouse()
	}
}

func (b *standardBank) BuyHotel(debtor Debtor, rental Buildable){
	if b.hotels > 0 {
		debtor.Pay(b, rental.BuildingPrice())
		b.hotels = b.hotels - 1
		rental.BuildHotel()
	}
}

func (b *standardBank) SellHouse(collector Collector, rental Buildable){
	rental.DemolishHouse()
	collector.Receive(rental.BuildingPrice() / 2, nil, nil)
	b.houses = b.houses + 1
}

func (b *standardBank) SellHotel(collector Collector, rental Buildable){
	rental.DemolishHotel()
	collector.Receive(rental.BuildingPrice()/2, nil, nil)
	b.hotels = b.hotels+ 1
}

func (b *standardBank) SellCard(collector Collector, card ConsumableCard){
	collector.Receive(card.Value(), nil, nil)
	b.cash = b.cash - card.Value()
}

func (b *standardBank) Mortgage(collector Collector, rental Rental) {
	if !rental.Mortgaged() {
		rental.Mortgage()
		collector.Receive(rental.MortgageValue(), nil, nil)
	}
}
func (b *standardBank) Unmortgage(debtor Debtor, rental Rental) {
	if rental.Mortgaged() {
		//TODO: Could lead to bankruptcy to bank if called on accidentally or unnecessarily
		debtor.Pay(b, rental.MortgageValue() * 1.2)
		rental.Mortgage()
	}
}