package models

//Foreclosure can be foreclosed
type Foreclosure interface {
	Foreclose()
}

type Group string

const (
	GREEN = Group("GREEN")
	PURPLE= Group("PURPLE")
	ORANGE= Group("ORANGE")
	RED= Group("RED")
	YELLOW= Group("YELLOW")
	BLUE= Group("BLUE")
	LIGHT_BLUE= Group("LIGHT_BLUE")
	BROWN= Group("BROWN")
	RAILROAD= Group("RAILROAD")
	UTILITY= Group("UTILITY")

)

type Rent struct {
	rent map[int32]int64
}
func (r *Rent) Rent(index int32) int64{
	return r.rent[index]
}


type Rental interface {
	Name() string
	Group() Group
	Landlord() Landlord
	Buy(Landlord) error
	Mortgage() error
	Unmortgage() error
	Mortgaged() bool
	Price() int64
	Monopoly() bool
	Rent() int64

}

type Buildable interface {
	SellBuilding() int64
	BuyBuilding() int64
}



