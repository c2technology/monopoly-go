package player

import (
	"fmt"
	"github.com/c2technology/monopoly-go/lib/game"
	"log"
	"math/rand"
)

//Creates a new Player
func NewPlayer(banker game.Banker) game.Player {
	player := &defaultPlayer{
		name:   playerNames[rand.Intn(len(playerNames))],
		banker: banker,
		//TODO: Get random strategy
	}
	log.Printf("%s has taken a seat!", player)
	return player
}

var playerNames = []string{
	"James",
	"John",
	"Robert",
	"Michael",
	"William",
	"David",
	"Richard",
	"Joseph",
	"Thomas",
	"Charles",
	"Christopher",
	"Daniel",
	"Matthew",
	"Anthony",
	"Donald",
	"Mark",
	"Paul",
	"Steven",
	"Andrew",
	"Kenneth",
	"George",
	"Joshua",
	"Kevin",
	"Brian",
	"Edward",
	"Ronald",
	"Timothy",
	"Jason",
	"Jeffrey",
	"Ryan",
	"Gary",
	"Jacob",
	"Nicholas",
	"Eric",
	"Stephen",
	"Jonathan",
	"Larry",
	"Justin",
	"Scott",
	"Frank",
	"Brandon",
	"Raymond",
	"Gregory",
	"Benjamin",
	"Samuel",
	"Patrick",
	"Alexander",
	"Jack",
	"Dennis",
	"Jerry",
	"Tyler",
	"Aaron",
	"Henry",
	"Douglas",
	"Jose",
	"Peter",
	"Adam",
	"Zachary",
	"Nathan",
	"Walter",
	"Harold",
	"Kyle",
	"Carl",
	"Arthur",
	"Gerald",
	"Roger",
	"Keith",
	"Jeremy",
	"Terry",
	"Lawrence",
	"Sean",
	"Christian",
	"Albert",
	"Joe",
	"Ethan",
	"Austin",
	"Jesse",
	"Willie",
	"Billy",
	"Bryan",
	"Bruce",
	"Jordan",
	"Ralph",
	"Roy",
	"Noah",
	"Dylan",
	"Eugene",
	"Wayne",
	"Alan",
	"Juan",
	"Louis",
	"Russell",
	"Gabriel",
	"Randy",
	"Philip",
	"Harry",
	"Vincent",
	"Bobby",
	"Johnny",
	"Logan",
	"Mary",
	"Patricia",
	"Jennifer",
	"Elizabeth",
	"Linda",
	"Barbara",
	"Susan",
	"Jessica",
	"Margaret",
	"Sarah",
	"Karen",
	"Nancy",
	"Betty",
	"Lisa",
	"Dorothy",
	"Sandra",
	"Ashley",
	"Kimberly",
	"Donna",
	"Carol",
	"Michelle",
	"Emily",
	"Amanda",
	"Helen",
	"Melissa",
	"Deborah",
	"Stephanie",
	"Laura",
	"Rebecca",
	"Sharon",
	"Cynthia",
	"Kathleen",
	"Amy",
	"Shirley",
	"Anna",
	"Angela",
	"Ruth",
	"Brenda",
	"Pamela",
	"Nicole",
	"Katherine",
	"Virginia",
	"Catherine",
	"Christine",
	"Samantha",
	"Debra",
	"Janet",
	"Rachel",
	"Carolyn",
	"Emma",
	"Maria",
	"Heather",
	"Diane",
	"Julie",
	"Joyce",
	"Evelyn",
	"Frances",
	"Joan",
	"Christina",
	"Kelly",
	"Victoria",
	"Lauren",
	"Martha",
	"Judith",
	"Cheryl",
	"Megan",
	"Andrea",
	"Ann",
	"Alice",
	"Jean",
	"Doris",
	"Jacqueline",
	"Kathryn",
	"Hannah",
	"Olivia",
	"Gloria",
	"Marie",
	"Teresa",
	"Sara",
	"Janice",
	"Julia",
	"Grace",
	"Judy",
	"Theresa",
	"Rose",
	"Beverly",
	"Denise",
	"Marilyn",
	"Amber",
	"Madison",
	"Danielle",
	"Brittany",
	"Diana",
	"Abigail",
	"Jane",
	"Natalie",
	"Lori",
	"Tiffany",
	"Alexis",
}

//Player is a member of the game
type defaultPlayer struct {
	banker  game.Banker
	cash    int64
	name    string
	rentals []game.Rental
	cards   []game.ConsumableCard
	//	piece GamePiece
	inJail        bool
	jailRemaining int
	//	gm GameMaster
}

func (p *defaultPlayer) Pay(collector game.Collector, owed int64) {
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
		//liquidate card
		for _, cardHeld := range p.cards {
			p.banker.SellCard(p, cardHeld)
			p.remove(p.cards, cardHeld)
		}
		//mortgage properties
		groups := make(map[game.Group]interface{})
		for _, rental := range p.rentals {
			groups[rental.Group()] = nil
			p.banker.Mortgage(p, rental)
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
			p.banker.Mortgage(p, rental)
			if owed <= p.cash {
				return
			}
		}
	}
}

func (p *defaultPlayer) Receive(cash int64, rentals []game.Rental, cards []game.ConsumableCard) {
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

func (p *defaultPlayer) WantToBuy(rental game.Rental) bool {
	//TODO: Implement BuyerStrategy
	return true
}

func (p *defaultPlayer) InJail() bool {
	return p.inJail
}

func (p *defaultPlayer) SentenceRemaining() int {
	return p.jailRemaining
}
func (p *defaultPlayer) PostBail() {
	//TODO

}

func (p *defaultPlayer) Build(group game.Group) {
	//TODO: Build evenly

}

func (p *defaultPlayer) liquidateBuilding(group game.Group) bool {
	var groupRentals []game.Rental

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

	var sortedRentals []game.Rental
	maxBuildings := int(0)
	for _, rental := range groupRentals {
		buildable, ok := rental.(game.Buildable)
		if ok {
			buildings := buildable.Houses() + buildable.Hotels()
			if buildings > maxBuildings {
				maxBuildings = buildings
				sortedRentals[maxBuildings] = rental
			}
		}
	}

	//Sell a building from a rental with the most amount of buildings
	buildable, ok := sortedRentals[maxBuildings].(game.Buildable)
	if ok {
		if buildable.Hotels() > 0 {
			p.banker.SellHotel(p, buildable)
		} else {
			p.banker.SellHouse(p, buildable)
		}
	}
	return true
}

func (p *defaultPlayer) declareBankruptcy(collector game.Collector) {
	collector.Receive(p.cash, p.rentals, p.cards)
	p.cash = 0
	p.rentals = nil
	p.cards = nil
	//TODO: gm.Eject(p) <-- gm should have the only reference to this player

}

func (p *defaultPlayer) String() string {
	return fmt.Sprintf("%s ($%v)", p.name, p.cash/100)
}

func (p *defaultPlayer) remove(elements []game.ConsumableCard, elem game.ConsumableCard) []game.ConsumableCard {
	index := 0
	for key, element := range elements {
		if element == elem {
			index = key
			break
		}
	}
	elements[len(elements)-1], elements[index] = elements[index], elements[len(elements)-1]
	elements = elements[:len(elements)-1]
	return elements
}
