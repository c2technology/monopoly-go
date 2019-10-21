# Monopoly Go

## Domain Driven Design

This project represents a learning playground for Domain Drive design within a Functional Programming context. The domain for this project being the popular board game Monopoly.

## Goals

I will be leveraging the rules of Monopoly to derive Domain Elements defining Services, Entities, and Value Objects. I will also identify the bounded contexts in which smaller models will operate. For each of the Bounded Contexts, I will identify the states in which they move through Monopoly. Each Bounded Context, both self-contained and comprised of smaller self-contained models, will flow through these defined states during the course of a game.

I will focus on immutability as a design construct to keep functions pure and to see the effect of this immutability on the complier's optimizations and memory allocation.

## Domain Model Elements

The services defined for this project will relate to the domain models within the game. Namely, the core services will be Banking, Community Chest, Chance, Dice, and the actual game Board. These services will allow interactions between Entities and Value Objects and contain the core logic for their respective domains.

### Entities

An Entity is an object that has an Identity and a Lifecycle. The only Entities in Monopoly are Players, and Properties. A Player lifecycle (below) is defined as "New" -> "Active" -> "Bankrupt" while a Property lifecycle is "Unowned" -> "Owned" -> "Mortgaged"/"Unmortgaged"
 
 #### Player
 A `Player` represents an individual participant in the game. The `Player` Entity will hold domain information related to Entities owned by the `Player`. These Entities include any `Property` owned by the `Player` as well as available cash, and held `Cards` (a "Get Out of Jail Free" card for example).

 ##### Lifecycle
 A Player lifecycle begins when they join a game. Each Player starts with $1500 in cash. They each start on the Go space. Priority order is determined by random draw. This is the starting state of a "new" Player.
  
  Throughout the course of the Game, a Player is considered Active. During this phase of the Player Lifecycle, the Player is buying properties, paying rent and fees, building houses, and collecting cash. This is the meat of the Player Lifecycle.
  
  The final state of a Player is Bankrupt. When a Player cannot afford to pay cash to either another Player or to the Bank, the Player becomes Bankrupt. When determining if a Player becomes Bankrupt, they must first sell all items of value (houses and hotels at half-price, and "Get Out of Jail Free" cards at $50), and mortgage properties. If the Player cannot afford to pay their debt with the cash accumulated in this way, they must turn over all of their assets to the creditor. If the creditor is the Bank, all Properties are immediately auctioned. If the creditor is another Player, that Player receives all Property and cash from the Bankrupt Player. The Creditor must then choose to either pay just the interest (10%) on or completely unmortgage any Property gained in this way. If a creditor decides to only pay interest (10%), they must pay the interest again when they decide to unmortgage that property.
    
 #### Property
  A `Property` holds defining attributes of a given individual property including its name, group, rent, housing costs, mortgage value, and mortgage state.
   
  1. Mediterranean Avenue
     1. Rent
         1. No House:
         1. 1 House:
         1. 2 Houses:
         1. 3 Houses:
         1. 4 Houses:
         1. Hotel:
     1. House Cost:
     1. Mortgage Value:
  1. Baltic Avenue
     1. Rent
          1. No House:
          1. 1 House:
          1. 2 Houses:
          1. 3 Houses:
          1. 4 Houses:
          1. Hotel:
      1. House Cost:
      1. Mortgage Value:
  1. Reading Railroad
     1. Rent
          1. 1 Railroad:
          1. 2 Railroad:
          1. 3 Railroad:
          1. 4 Railroad:
      1. Mortgage Value:
  1. Oriental Avenue
     1. Rent
          1. No House:
          1. 1 House:
          1. 2 Houses:
          1. 3 Houses:
          1. 4 Houses:
          1. Hotel:
     1. House Cost:
     1. Mortgage Value:
  1. Vermont Avenue
     1. Rent
          1. No House:
          1. 1 House:
          1. 2 Houses:
          1. 3 Houses:
          1. 4 Houses:
          1. Hotel:
     1. House Cost:
     1. Mortgage Value:
  1. Connecticut Avenue
     1. Rent
          1. No House:
          1. 1 House:
          1. 2 Houses:
          1. 3 Houses:
          1. 4 Houses:
          1. Hotel:
     1. House Cost:
     1. Mortgage Value:
  1. St. Charles Place
     1. Rent
          1. No House:
          1. 1 House:
          1. 2 Houses:
          1. 3 Houses:
          1. 4 Houses:
          1. Hotel:
     1. House Cost:
     1. Mortgage Value:
  1. Electric Company
     1. Rent
          1. 1 Utility:
          1. 2 Utilities:
      1. Mortgage Value:
  1. States Avenue
     1. Rent
          1. No House:
          1. 1 House:
          1. 2 Houses:
          1. 3 Houses:
          1. 4 Houses:
          1. Hotel:
     1. House Cost:
     1. Mortgage Value:
  1. Virginia Avenue
     1. Rent
         1. No House:
         1. 1 House:
         1. 2 Houses:
         1. 3 Houses:
         1. 4 Houses:
         1. Hotel:
      1. House Cost:
      1. Mortgage Value:
  1. Pennsylvania Railroad
      1. Rent
           1. 1 Railroad:
           1. 2 Railroad:
           1. 3 Railroad:
           1. 4 Railroad:
       1. Mortgage Value:
 1. St. James Place
     1. Rent
          1. No House:
          1. 1 House:
          1. 2 Houses:
          1. 3 Houses:
          1. 4 Houses:
          1. Hotel:
      1. House Cost:
      1. Mortgage Value:
  1. Tennessee Avenue
     1. Rent
          1. No House:
          1. 1 House:
          1. 2 Houses:
          1. 3 Houses:
          1. 4 Houses:
          1. Hotel:
      1. House Cost:
      1. Mortgage Value:
  1. New York Avenue
     1. Rent
          1. No House:
          1. 1 House:
          1. 2 Houses:
          1. 3 Houses:
          1. 4 Houses:
          1. Hotel:
      1. House Cost:
      1. Mortgage Value:
  1. Kentucky Avenue
     1. Rent
          1. No House:
          1. 1 House:
          1. 2 Houses:
          1. 3 Houses:
          1. 4 Houses:
          1. Hotel:
      1. House Cost:
      1. Mortgage Value:
  1. Indiana Avenue
     1. Rent
          1. No House:
          1. 1 House:
          1. 2 Houses:
          1. 3 Houses:
          1. 4 Houses:
          1. Hotel:
      1. House Cost:
      1. Mortgage Value:
  1. Illinois Avenue
     1. Rent
          1. No House:
          1. 1 House:
          1. 2 Houses:
          1. 3 Houses:
          1. 4 Houses:
          1. Hotel:
      1. House Cost:
      1. Mortgage Value:
  1. B & O Railroad
       1. Rent
            1. 1 Railroad:
            1. 2 Railroad:
            1. 3 Railroad:
            1. 4 Railroad:
        1. Mortgage Value:
  1. Atlantic Avenue
     1. Rent
          1. No House:
          1. 1 House:
          1. 2 Houses:
          1. 3 Houses:
          1. 4 Houses:
          1. Hotel:
      1. House Cost:
      1. Mortgage Value:
  1. Ventnor Avenue
     1. Rent
          1. No House:
          1. 1 House:
          1. 2 Houses:
          1. 3 Houses:
          1. 4 Houses:
          1. Hotel:
      1. House Cost:
      1. Mortgage Value:
  1. Water Works
     1. Rent
          1. 1 Utility:
          1. 2 Utilities:
      1. House Cost:
      1. Mortgage Value:
  1. Marvin Gardens
      1. Rent
          1. No House:
          1. 1 House:
          1. 2 Houses:
          1. 3 Houses:
          1. 4 Houses:
          1. Hotel:
      1. House Cost:
      1. Mortgage Value:
  1. Pacific Avenue
     1. Rent
          1. No House:
          1. 1 House:
          1. 2 Houses:
          1. 3 Houses:
          1. 4 Houses:
          1. Hotel:
      1. House Cost:
      1. Mortgage Value:
  1. North Carolina Avenue
     1. Rent
          1. No House:
          1. 1 House:
          1. 2 Houses:
          1. 3 Houses:
          1. 4 Houses:
          1. Hotel:
      1. House Cost:
      1. Mortgage Value:
  1. Pennsylvania Avenue
     1. Rent
          1. No House:
          1. 1 House:
          1. 2 Houses:
          1. 3 Houses:
          1. 4 Houses:
          1. Hotel:
      1. House Cost:
      1. Mortgage Value:
  1. Short Line
       1. Rent
            1. 1 Railroad:
            1. 2 Railroad:
            1. 3 Railroad:
            1. 4 Railroad:
        1. Mortgage Value:
  1. Park Place
     1. Rent
          1. No House:
          1. 1 House:
          1. 2 Houses:
          1. 3 Houses:
          1. 4 Houses:
          1. Hotel:
      1. House Cost:
      1. Mortgage Value:
  1. Boardwalk
     1. Rent
          1. No House:
          1. 1 House:
          1. 2 Houses:
          1. 3 Houses:
          1. 4 Houses:
          1. Hotel:
      1. House Cost:
      1. Mortgage Value:
  
### Value Objects
Each Entity given an actual value is considered a Value Object. Each time a Value Object would mutate, instead, a copy of the Value Object will be created that contains the updated value. While this is core concept of immutability, the Go complier will hopefully also optimize memory allocation and garbage collection as a single stack frame will contain both the old value (which will be garbage collected) and the new value that is returned (which will also be garbage collected, as the returned value in Go is copied to the calling stack frame).

#### Die
 Each Die that is used during the game will have a number of faces representing a specific value. This could be implemented in may ways (such as a speed Die) that would affect rules. For the purpose of this project, we will be using a simple 6 sided `Die`. 

#### Card 
 Each card has a title and a function determining the action that is taken when drawn. In most cases, the action is performed without `Player` interaction. However, there is a "Get Out of Jail Free" card in the Chance `Deck` and the Community Chest `Deck` that a `Player` may sell to the `Bank` for $50 or keep. This project will assume the `Player` will keep the `Card` until it is needed and will sell the `Card` when cash is needed.

 #### Space
 Each `Space` on the game board has some sort of event associated with it when a `Player` lands on or passes through the `Space`. These actions could be:
  1. The Space holds an unowned Property and is available for purchase by the Player.
  1. The Space holds an owned Property and the Player must pay rent to the owner.
  1. The Space is Chance or Community Chest and the Player must draw a Card from the respective Deck.
  1. The Space is a Luxury Tax and the Player must pay the Bank an amount of cash.
  1. The Space is Go and the Bank must pay the Player $200.
  1. The Space is Income Tax and the Player must choose to pay 10% of their assets or $200.
  1. The Space is Go To Jail and the Player moves directly to Jail.
  1. The Space is In Jail/Just Visiting and, depending on if the Player is currently on In Jail or Just Visiting, prevents the Player from moving or takes no action against the Player.
  1. The Space is Free Parking and nothing happens.
  
  The `Spaces` available in this project are:
    1. Go
        a. Action: Collect $200 from `Bank` when landing on or passing. 
    1. Mediterranean Avenue
        a. Action: If unowned, buy from `Bank` for $60 or auction. If owned, pay rent
    1. Community Chest
        a. Action: Draw `Card` from Community Chest `Deck`
    1. Baltic Avenue
        a. Action: If unowned, buy from `Bank` for $60 or auction. If owned, pay rent
    1. Income Tax
        a: Pay 10% of assets or $200
    1. Reading Railroad
        a. Action: If unowned, buy from `Bank` for $200 or auction. If owned, pay rent
    1. Oriental Avenue
        a. Action: If unowned, buy from `Bank` for $100 or auction. If owned, pay rent
    1. Chance
        a: Action: Draw `Card` from Chance `Deck`
    1. Vermont Avenue
        a Action: If unowned, buy from `Bank` for $100 or auction. If owned, pay rent
    1. Connecticut Avenue
        a. Action: If unowned, buy from `Bank` for $120 or auction. If owned, pay rent
    1 Just Visiting/In Jail
        a. Just Visiting Action: None
        b. In Jail Action: If a roll of doubles is made, `Player` may get out of Jail for free, forfeiting the extra roll. If 3 rolls have been made without doubles, the `Player` must pay $50 and move to Just Visiting. Alternatively, a `Player` may use a "Get Out of Jail Free" `Card` in lieu of paying $50 or rolling. 
    1. St. Charles Place
        a. Action: If unowned, buy from `Bank` for $140 or auction. If owned, pay rent
    1. Electric Company
        a. Action: If unowned, buy from `Bank` for $150 or auction. If owned, pay rent
    1. States Avenue
        a. Action: If unowned, buy from `Bank` for $140 or auction. If owned, pay rent
    1. Virginia Avenue
        a. Action: If unowned, buy from `Bank` for $160 or auction. If owned, pay rent
    1. Pennsylvania Railroad
        a. Action: If unowned, buy from `Bank` for $200 or auction. If owned, pay rent
    1. St. James Place
        a. Action: If unowned, buy from `Bank` for $180 or auction. If owned, pay rent
    1. Community Chest
        a. Action: Draw `Card` from Community Chest `Deck`
    1. Tennessee Avenue
        a. Action: If unowned, buy from `Bank` for $180 or auction. If owned, pay rent
    1. New York Avenue
        a. Action: If unowned, buy from `Bank` for $200 or auction. If owned, pay rent
    1. Free Parking
        a. Action: None
    1. Kentucky Avenue
        a. Action: If unowned, buy from `Bank` for $220 or auction. If owned, pay rent
    1. Chance
        a. Action: Draw `Card` from Chance `Deck`
    1. Indiana Avenue
        a. Action: If unowned, buy from `Bank` for $220 or auction. If owned, pay rent
    1. Illinois Avenue
        a. Action: If unowned, buy from `Bank` for $240 or auction. If owned, pay rent
    1. B & O Railroad
        a. Action: If unowned, buy from `Bank` for $200 or auction. If owned, pay rent
    1. Atlantic Avenue
        a. Action: If unowned, buy from `Bank` for $260 or auction. If owned, pay rent
    1. Ventnor Avenue
        a. Action: If unowned, buy from `Bank` for $260 or auction. If owned, pay rent
    1. Water Works
        a. Action: If unowned, buy from `Bank` for $280 or auction. If owned, pay rent
    1. Marvin Gardens
        a. Action: If unowned, buy from `Bank` for $300 or auction. If owned, pay rent
    1. Pacific Avenue
        a. Action: If unowned, buy from `Bank` for $300 or auction. If owned, pay rent
    1. North Carolina Avenue
        a. Action: If unowned, buy from `Bank` for $300 or auction. If owned, pay rent
    1. Community Chest
        a. Action: Draw `Card` from Community Chest `Deck`
    1. Pennsylvania Avenue
        a. Action: If unowned, buy from `Bank` for $320 or auction. If owned, pay rent
    1. Short Line
        a. Action: If unowned, buy from `Bank` for $200 or auction. If owned, pay rent
    1. Chance
        a. Action: Draw `Card` from Chance `Deck`
    1. Park Place
        a. Action: If unowned, buy from `Bank` for $350 or auction. If owned, pay rent
    1. Luxury Tax
        a. Action: Pay `Bank` $100.
    1. Boardwalk
        a. Action: If unowned, buy from `Bank` for $400 or auction. If owned, pay rent

#### Strategy
Each non-human `Player` will choose randomly from a pool of `Strategies`. These will determine how the `Player` makes decisions throughout the game. Strategies used during this project are:
1. Aggressive 
    a. Purchases: The `Player` will attempt to purchase, at retail price, all `Property` that is available, mortgaging any `Property` that is not part of a full set.
    b. Auction: The `Player` will attempt to purchase, up to retail price, all auctioned `Property`, mortgaging any `Property` that is not part of a full set.
    c. Housing: The `Player` will attempt to purchase as many houses as possible when a full property is acquired, mortgaging any `Property` that is not part of a full set.
    d. Unmortgaging: The `Player` will attempt to unmortgage as many `Properties` as possible if there is no available `Property` or housing available to purchase. Full set `Property` is prioritized, then highest rental value.
    e. Trading: Trades are outside the scope of this project.
1. Conservative 
    a. Purchases: The `Player` will attempt to purchase, at retail price, all `Property` that is available.
    b. Auction: The `Player` will attempt to purchase, up to retail price, all auctioned `Property` without mortgaging any `Property`.
    c. Housing: The `Player` will attempt to purchase as many houses as possible when a full property is acquired without mortgaging any `Property`.
    d. Unmortgaging: The `Player` will attempt to unmortgage as many `Properties` as possible if there is no available `Property` or housing available to purchase. Full set `Property` is prioritized, then highest rental value.
    e. Trading: Trades are outside the scope of this project.
 
### Bounded Contexts
Bounded Contexts are areas in which smaller Domain Models define nontrivial domain logic using a collection of Entities operating one complete area of functionality. Each Bounded Context is considered a smaller subset of the entire Domain Model and is responsible for a subset of Domain Model logic. Bounded Contexts are themselves self-contained modules representing a complete areas of functionality within a system and together make the entire Domain Model.

The Bounded Contexts I will be using for this project are:

#### Deck
 The Deck holds an ordered collection of Card Value Objects and knows how to shuffle and deal cards. This represents the complete functionality of Chance and Community Chest decks. 
 
 For cross-boundary interactions, such as drawing a card, a Service would be used to determine if and when a Player draws a card.

#### Bank
 The Bank handles the Entities related to the bank in Monopoly. This includes cash transactions, property purchases and auctions, house and hotel purchases, and bankruptcy.
 
 A service would determine when the Bank would be invoked for a particular cross-boundary action such as assessing a fee against a `Player`, starting an auction, invoking the sale of a `Property`..

#### Dice
 The Dice use both Die Value Objects to determine the values rolled and if the value is a double. 
 
 Services that interact with the Dice would determine what happens when multiple doubles are rolled.

#### Board
 A Board holds the individual Space and layout on the board. It also holds the Player locations.
 
 A Service would determine which `Player` has priority during the turn. All other turn related actions are handled by the `Board`. 

### Services
 Services facilitate cross Boundary Context interactions. Services are typically limited in a Functional Domain Model and in this project, there is only one. 
 
#### Game

The `Game` service serves as the controlling force across the entire Domain Model. It determines the `Player` order, the starting cash amounts for each `Player`, the roll of the `Dice`, where a `Player` moves, how a `Player` moves, any action related to a `Space` and any action related to a `Card`. The `Game` service is the essential orchestrator of the Domain Model.

## Pure Functions

In contrast to the "go fast" paradigm of the Go language, I will be leveraging immutability of code in this project to attempt to make each function within a given domain model _pure_. That is, given exactly one input, each function will produce exactly one output. In reaching this purity, each function will only rely on the input given to derive an output. For any given function, there should be no external dependencies. Any requirement or dependency to a function will be provided as an input.  
