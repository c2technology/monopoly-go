package classic

import (
	"github.com/c2technology/monopoly-go/lib/game"
	"log"
	"github.com/c2technology/monopoly-go/lib/card"
)

func NewCommunityChest(banker game.Banker) game.Deck {
	log.Println("Initializing Community Chest...")

	var cards []game.Card

	cards = append(cards, card.NewCard("Advance to Go (Collect $200)", func(player game.Player) {
		//game.Move(player, 0)
		banker.Pay(player, 200)
	}))
	cards = append(cards, card.NewCard("Bank error in your favor – Collect $200", func(player game.Player) {
		banker.Pay(player, 200)
	}))
	cards = append(cards, card.NewCard("Doctor's fee – Pay $50", func(player game.Player) {
		player.Pay(banker, 150)
	}))

	cards = append(cards, card.NewCard("From sale of stock you get $45", func(player game.Player) {
		banker.Pay(player, 45)
	}))
	cards = append(cards, card.NewConsumableCard("Get out of Jail, Free – This card may be kept until needed or sold", func(player game.Player) {
		//put card in player inventory
	}, func(player game.Player) {
		//Use card
	}, 50))
	cards = append(cards, card.NewCard("Go to Jail – Go directly to jail – Do not pass Go – Do not collect $200", func(player game.Player) {
		//space := game.Move(player, space.Jail)
	}))

	cards = append(cards, card.NewCard("Grand Opening – Collect $50 from every player", func(player game.Player) {
		//for competitor : game.Players() {
		//	competitor.Pay(player, 50)
		//}
	}))
	cards = append(cards, card.NewCard("Holiday Fund matures - Collect $100", func(player game.Player) {
		banker.Pay(player, 100)
	}))
	cards = append(cards, card.NewCard("Income tax refund – Collect $20", func(player game.Player) {
		banker.Pay(player, 20)
	}))
	cards = append(cards, card.NewCard("Life insurance matures – Collect $100", func(player game.Player) {
		banker.Pay(player, 100)
	}))

	cards = append(cards, card.NewCard("Pay hospital fees of $100 {Pay hospital $100}", func(player game.Player) {
		banker.Pay(player, 100)
	}))

	cards = append(cards, card.NewCard("Go Back 3 Spaces.", func(player game.Player) {
		//space := game.Move(player, -3)
		//space.DoAction(player)
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
	return card.NewBasicDeck("Chance", cards)
}
