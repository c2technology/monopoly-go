package models


//Creates a new Rent structured for housing properties
func newHousingRent(base, house1, house2, house3, house4, hotel int64) *Rent {
	rent := make(map[int32]int64)
	rent[0] = base
	rent[1] = house1
	rent[2] = house2
	rent[3] = house3
	rent[4] = house4
	rent[5] = hotel
	return &Rent{rent}
}

//Creates a new Rent structured for utilities
func newUtilityRent(multiplier1, multiplier2 int64) *Rent {
	rent := make(map[int32]int64)
	rent[0] = multiplier1
	rent[1] = multiplier2
	return &Rent{rent}
}

//Creates a new Rent structured for railroads
func newRailroadRent(tier1, tier2, tier3, tier4 int64) *Rent {
	rent := make(map[int32]int64)
	rent[0] = tier1
	rent[1] = tier2
	rent[2] = tier3
	rent[3] = tier4
	return &Rent{rent}
}

//Creates a new Rental
func NewRentalProperty(name string, group Group, rent Rent, mortgage, price, buildingCost int64) Rental {
	return &RentalProperty{
		mortgage:     mortgage,
		price:        price,
		name:         name,
		rent:         rent,
		group:        group
	}
}
