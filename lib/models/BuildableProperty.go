package models

import "errors"

//A Rental Property with Buildings
type BuildingProperty struct{
	*rentalProperty
	buildings    int32
	buildingCost int64
}

//Determines rent with buildings
func (r *BuildingProperty) Rent() int64 {
	if r.buildings < 1 {
		return r.rentalProperty.Rent()
	}
	return r.rent.Rent(r.buildings)
}

//Increases rent if possible
func (r *BuildingProperty) increaseRent(){
	if r.buildings < 5 {
		r.buildings = r.buildings + 1
	}
}

//Decreases rent if possible
func (r *BuildingProperty) decreaseRent(){
	if r.hasBuildings() {
		r.buildings = r.buildings - 1
	}
}

//Does this property have buildings?
func (r *BuildingProperty) hasBuildings() bool {
	return r.buildings > 0
}

//Attempts to sell a building for an amount returned
func (r *BuildingProperty) SellBuilding() int64 {
	//TODO: Need to evenly sell
	if r.buildings > 0 {
		r.decreaseRent()
		//logger.debug(landlord.toString() + " sold a building on " + this.toString());
		return r.buildingCost / 2;
	}
	return 0
}

//Attempts to buy a building for amount returned
func (r *BuildingProperty) BuyBuilding() int64{
	//TODO: Need to evenly buy
	if r.buildings < 5 {
		r.increaseRent()
		//logger.debug(landlord.toString() + " bought a building on " + this.toString());
		return r.buildingCost
	}
	return 0
}

