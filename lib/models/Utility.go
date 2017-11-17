package models

import "errors"

type Utility struct{
	landlord  Landlord
	mortgaged bool
	mortgage  int64
	price     int64
	name      string
	monopoly  bool
	rent      Rent
	color     Group
}

//Forecloses the Rental
func (r *Utility) Foreclose() {
	r.mortgaged = false
	r.monopoly= false
	r.landlord= nil
	r.buildings = 0
}

//Name of the place
func (r *Utility) Name() string {
	return r.name
}

//Group group of the Rental
func (r *Utility) Color() Group {
	return r.color
}

//Is this Rental currently owned?
func (r *Utility) Landlord() Landlord {
	return r.landlord
}

//Sets the new Landlord of the Rental
func (r *Utility) Buy(bidder Landlord) error {
	if err := bidder.Spend(r.price, nil, nil); err!= nil {
		return err
	}
	if r.landlord != nil {
		r.landlord.Receive(r.price, nil, nil)
	}
	//TODO: bank receives funds (although this could be rectified by never having a nil owner and the bank owns everything by default)
	r.landlord=bidder;
	return nil
}

//Mortgages the Rental
func (r *Utility) Mortgage() error {
	if r.landlord != nil && !r.mortgaged {
		r.landlord.Receive(r.mortgage, nil, nil)
		r.mortgaged = true
		return nil
	}
	return errors.New("property unowned or already mortgaged")
}

//Sets the Mortgaged state of the Rental
func (r *Utility) Unmortgage() error {
	if r.landlord != nil && r.mortgaged {
		r.landlord.Spend(r.mortgage * 1.2, nil, nil)
		r.mortgaged = false
		return nil
	}
	return errors.New("property unowned or not mortgaged")
}

//Mortgaged state of the Rental
func (r *Utility) Mortgaged() bool {
	return r.mortgaged
}

//Price of the Rental
func (r *Utility) Price() int64 {
	return r.price
}

//Determines if the Rental is part of a monopoly
func (r *Utility) Monopoly() bool {
	return r.monopoly
}

//Rent is doubled when a monopoly exists and their are no buildings. Otherwise, it's whatever the rent index is set to
func (r *Utility) Rent() int64{
	if r.buildings == 0 && r.monopoly {
		return r.rent.Rent(r.buildings) * 2
	}
	return r.rent.Rent(r.buildings)
}

//Increases rent if possible
func (r *Utility) increaseRent(){
	if r.buildings < 5 {
		r.buildings = r.buildings + 1
	}
}

//Decreases rent if possible
func (r *Utility) decreaseRent(){
	if r.buildings > 0 {
		r.buildings = r.buildings - 1
	}
}

//Mortgageable is no buildings, it is owned, and it's not already mortgaged
func (r *Utility) Mortgagable() bool{
	return !r.hasBuildings() && r.Owned() && !r.Mortgaged()
}

//Does this property have buildings?
func (r *Utility) hasBuildings() bool {
	return r.buildings > 0
}

//Attempts to sell a building for an amount returned
func (r *Utility) SellBuilding() int64 {
	//TODO: Need to evenly sell
	if r.buildings > 0 {
		r.decreaseRent()
		//logger.debug(landlord.toString() + " sold a building on " + this.toString());
		return r.buildingCost / 2;
	}
	return 0
}

//Attempts to buy a building for amount returned
func (r *Utility) BuyBuilding() int64{
	//TODO: Need to evenly buy
	if r.buildings < 5 {
		r.increaseRent()
		//logger.debug(landlord.toString() + " bought a building on " + this.toString());
		return r.buildingCost
	}
	return 0
}

