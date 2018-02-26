package classic

import (
	"github.com/c2technology/monopoly-go/lib/game"
	"log"
	"github.com/c2technology/monopoly-go/lib/card"
)

func NewChance(banker game.Banker) game.Deck {
	log.Println("Initializing Chance...")

	var cards []game.Card

	cards = append(cards, card.NewCard("Advance to Go", func(player game.Player) {
		banker.Pay(player, 200)
	}))
	cards = append(cards, card.NewCard("Advance to Illinois", func(player game.Player) {
		//game.Move(player, space)
	}))
	cards = append(cards, card.NewCard("Advance to St. Charles Place", func(player game.Player) {
		//game.Move(player, space)
	}))

	cards = append(cards, card.NewCard("Advance token to the nearest Railroad and pay owner twice the rental to which he/she {he} is otherwise entitled. If Railroad is unowned, you may buy it from the Bank.", func(player game.Player) {
		//space := game.Move(player, space)
		//if space.Owned() {
		//	player.Pay(space.Owner(), space.Rent() * 2)
		//}
	}))
	cards = append(cards, card.NewCard("Advance token to the nearest Railroad and pay owner twice the rental to which he/she {he} is otherwise entitled. If Railroad is unowned, you may buy it from the Bank.", func(player game.Player) {
		//space := game.Move(player, space)
		//if space.Owned() {
		//	player.Pay(space.Owner(), space.Rent() * 2)
		//}
	}))
	cards = append(cards, card.NewCard("Advance token to nearest Utility. If unowned, you may buy it from the Bank. If owned, throw dice and pay owner a total ten times the amount thrown", func(player game.Player) {

	}))
	cards = append(cards, card.NewCard("Bank pays you dividend of $50", func(player game.Player) {
		banker.Pay(player, 50)
	}))

	cards = append(cards, card.NewConsumableCard("Get out of Jail Free – This card may be kept until needed, or traded/sold", func(player game.Player) {
		//put card in player inventory
	}, func(player game.Player) {
		//Use card
	}, 50))

	cards = append(cards, card.NewCard("Go Back 3 Spaces.", func(player game.Player) {
		//space := game.Move(player, -3)
		//space.DoAction(player)
	}))
	cards = append(cards, card.NewCard("Go to Jail – Go directly to Jail – Do not pass Go, do not collect $200", func(player game.Player) {
		//space := game.Move(player, space.Jail)
	}))
	cards = append(cards, card.NewCard("Make general repairs on all your property – For each house pay $25 – For each hotel $100", func(player game.Player) {
		//for buildable : player.getBuldables(){
		//	if (buildable.Hotels() > 0 ){
		//		player.Pay(game.Banker(), 100 * buildable.Hotels()
		//	else {
		//		player.Pay(game.Banker(), 100 * buildable.Houses()
		//	}
		//}
	}))
	cards = append(cards, card.NewCard("Pay poor tax of $15", func(player game.Player) {
		player.Pay(banker, 15)
	}))
	cards = append(cards, card.NewCard("Take a ride on the Reading – If you pass Go, collect $200", func(player game.Player) {
		//space := game.Move(player, Reading)
		//Do some math to determine if Go was passed
		//space.DoAction(player)
	}))
	cards = append(cards, card.NewCard("You have been elected Chairman of the Board – Pay each player $50", func(player game.Player) {
		//players := game.Players()
		//if player.Cash() > 50 * game.Players()
		//	for recipient : players
		//		player.Pay(recipient, 50)
		//		player.Pay(recipient, 50)
		// 	}
		//} else {
		//	special case that no player is paid out if the player cannot pay all
		//	player.DeclareBankruptcy()
		//}
	}))
	cards = append(cards, card.NewCard("Your building {and} loan matures – Collect $150", func(player game.Player) {
		banker.Pay(player, 150)
	}))

	deck := &basicDeck{name: "Chance", cards: cards}
	deck.Shuffle()
	return deck

}
