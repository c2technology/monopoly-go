package game

//Collects assets
type Collector interface {
	//Receive assets as a one sided-exchange
	Receive(int64, []Rental, []ConsumableCard)
}
