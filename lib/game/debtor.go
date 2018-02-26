package game

//Someone who owes something
type Debtor interface {
	//Pays the given amount to the given Collector. Steps to ensure payment is made should be performed.
	Pay(Collector, int64)
}
