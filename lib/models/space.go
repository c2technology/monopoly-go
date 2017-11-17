package models



//Space is a space on the Board. It could be a Property or could be an Action
type Space interface {
	Name() string
	DoAction() error
}

//PropertyManager manages a collection of Rentals
type PropertyManager struct {
	rentals []Rental
}
func (pm *PropertyManager) Foreclosures() []Foreclosure {
	var foreclosures []Foreclosure
	for _, rental := range pm.rentals {
		foreclosure, ok := rental.(Foreclosure)
		if ok {
			foreclosures = append(foreclosures, foreclosure)
		}
	}
	return foreclosures
}

func (pm *PropertyManager) Properties() []*Property {
	var properties []*Property
	for _, rental := range pm.rentals {
		property, ok := rental.(*Property)
		if ok {
			properties = append(properties, property)
		}
	}
	return properties
}

func (pm *PropertyManager) Railroads() []*Railroad {
	var railroads []*Railroad
	for _, rental := range pm.rentals {
		railroad, ok := rental.(*Railroad)
		if ok {
			railroads = append(railroads, railroad)
		}
	}
	return railroads
}

func (pm *PropertyManager) Utilities() []*Utility {
	var utilities []*Utility
	for _, rental := range pm.rentals {
		utility, ok := rental.(*Utility)
		if ok {
			utilities = append(utilities, utility)
		}
	}
	return utilities
}
func (pm *PropertyManager) Buildable() []Buildable {
	var buildables []Buildable
	for _, rental := range pm.rentals {
		buildable, ok := rental.(Buildable)
		if ok {
			buildables = append(buildables, buildable)
		}
	}
	return buildables
}

func (pm *PropertyManager) Rentals(color Group) []*Rental {
	var rentals []*Rental
	for _, rental := range pm.rentals {
		if rental.Group() == color {
			rentals = append(rentals, rental)
		}
	}
	return rentals
}