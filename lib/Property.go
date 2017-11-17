package lib

//Property is any one of the various properties within the game
type Property interface {
	Name() string
	Rent() int64
	BuildingPrice() int64
	Price() int64
	Houses() int32
	Hotels() int32
	IsMortgaged() bool
	Mortgage() int64
}

//Player is a member of the game
type Player interface {
	Name() string
	Cash() int64
	Properties() []Property
	Actions() []Action
}



//Action represents Community Chest and Chance cards
type Action interface {
	Description() string
	DoAction() error
}



type Die interface {
	Roll() int32
}