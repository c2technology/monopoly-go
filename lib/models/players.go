package models


//Someone who owns something
type Proprietor interface {
	//Check to see if this Landlord owns the Rental
	HasRental(Rental) bool

	//Check to see if this Landlord owns the Card
	HasCard(Card) bool

	//Receives assets
	Receive(int64, []Rental, []Card)
}

//Someone who owes something
type Debtor interface {
	//How much cash is available
	Cash() int64

	//Spends assets. Returns error if unable to spend
	Spend(int64, []Rental, []Card) error
}

//Landlord collects assets
type Landlord interface {




}

//Renter pays
type Renter interface {
	//Declares bankruptcy and gives all assets to the Landlord
	DeclareBankruptcy(Landlord)
	//Pays a specific amount to the Landlord
	PayRent(Landlord, int64)
}

//CardCollector collects cards
type CardCollector interface {
	//Collects a CollectibleCard
	CollectCard(CollectibleCard)
	//Uses a CollectibleCard
	UseCard() bool
	//Sells a a CollectibleCard for an amount
	SellCards(int64)
}

//Player is a member of the game
type Player struct {
	cash int64
	name string
	properties PropertyManager
	cards []CollectibleCard
	piece GamePiece
	inJail bool
	jailRemaining int32
	gm GameMaster
}


func NewPlayer(name string, piece GamePiece) *Player{
	return &Player{name: name, piece: piece}
}