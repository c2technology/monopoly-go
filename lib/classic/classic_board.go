package classic

import (
	"github.com/c2technology/monopoly-go/lib/game"
	"github.com/c2technology/monopoly-go/lib/property"
	"github.com/c2technology/monopoly-go/lib/board"
)

type classicBoard struct {
	spaces map[int]game.Space
}

func (c *classicBoard) Space(player game.Player) game.Space {
	return nil
}

func (c *classicBoard) Move(player game.Player, location int) game.Space {
	return nil
}

func NewClassicBoard(gameState game.GameState) {
	spaces := make(map[int]game.Space)

	communityChest := board.NewActionableSpace("Community Chest", func(player game.Player){
		gameState.CommunityChest().Draw(player)
	})

	//chance := NewActionableSpace("Chance", func(player game.Player){
	//	gameState.Chance().Draw(player)
	//})

	//Go
	spaces[0] = board.NewActionableSpace("Go", func(player game.Player){
		gameState.Banker().Pay(player, 200)
	})

	//Mediterranean Avenue
	spaces[1] = property.NewRental("Mediterranean Avenue", game.PURPLE, gameState.Banker(), 60, 30, 50,map[int]int64{
		0: 2,
		1: 10,
		2: 30,
		3: 90,
		4: 180,
		5: 250,
	})

	//Comunity Chest
	spaces[2]= communityChest

	//Baltic Avenue
	spaces[3] = property.NewRental("Baltic Avenue", game.PURPLE, gameState.Banker(), 60, 30, 50,map[int]int64{
		0: 4,
		1: 20,
		2: 60,
		3: 180,
		4: 320,
		5: 450,
	})



	//Income Tax
	//Reading Railroad
	//TODO: Railroad property?
	//spaces[3] = property.NewRental("Reading Railroad", game.RAILROAD, gameState.Banker(), 200, 100, 0,map[int]int64{
	//	0: 25,
	//	1: 50,
	//	2: 100,
	//	3: 200,
	//})


	//Oriental Avenue
	//Chance
	//Vermont Avenue
	//Connecticut Avenue
	//Jail
	//St. Charles Place
	//Electric Company
	//States Avenue
	//Virginia Avenue
	//Pennsylvania Avenue
	//St. James Place
	//Community Chest
	//Tennesee Avenue
	//New York Avenue
	//Free Parking
	//Kentucky Avenue
	//Chance
	//Indiana Avenue
	//Illinois Avenue
	//B & O Railroad
	//Atlantic Avenue
	//Ventnor Avenue
	//Water Works
	//Marvin Gardens
	//Go To Jail
	//Pacific Avenue
	//North Carolina Avenue
	//Community Chest
	//Pennsylvania Avenue
	//Short Line
	//Chance
	//Park Place
	//Luxury Tax
	//Boardwalk




}
