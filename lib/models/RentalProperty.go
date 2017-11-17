package models

import "errors"

//Rental Property is any Property
type rentalProperty struct{
	landlord  Landlord
	mortgaged bool
	mortgage  int64
	price     int64
	name      string
	monopoly  bool
	rent      Rent
	group     Group
}

//Forecloses the Rental
func (r *rentalProperty) Foreclose() {
	r.mortgaged = false
	r.monopoly= false
	r.landlord= nil
}

//Name of the place
func (r *rentalProperty) Name() string {
	return r.name
}

//Group group of the Rental
func (r *rentalProperty) Color() Group {
	return r.group
}

//Is this Rental currently owned?
func (r *rentalProperty) Landlord() Landlord {
	return r.landlord
}

//Sets the new Landlord of the Rental
func (r *rentalProperty) Buy(bidder Landlord) error {
	if err := bidder.Spend(r.price, nil, nil); err!= nil {
		return err
	}
	if r.landlord != nil {
		r.landlord.Receive(r.price, nil, nil)
	}
	//TODO: bank receives funds (although this could be rectified by never having a nil owner and the bank owns everything by default)
	r.landlord=bidder;
	return nil
}

//Mortgages the Rental
func (r *rentalProperty) Mortgage() error {
	if r.landlord != nil && !r.mortgaged {
		r.landlord.Receive(r.mortgage, nil, nil)
		r.mortgaged = true
		return nil
	}
	return errors.New("property unowned or already mortgaged")
}

//Sets the Mortgaged state of the Rental
func (r *rentalProperty) Unmortgage() error {
	if r.landlord != nil && r.mortgaged {
		r.landlord.Spend(r.mortgage * 1.2, nil, nil)
		r.mortgaged = false
		return nil
	}
	return errors.New("property unowned or not mortgaged")
}

//Mortgaged state of the Rental
func (r *rentalProperty) Mortgaged() bool {
	return r.mortgaged
}

//Price of the Rental
func (r *rentalProperty) Price() int64 {
	return r.price
}

//Determines if the Rental is part of a monopoly
func (r *rentalProperty) Monopoly() bool {
	return r.monopoly
}

//Rent is doubled when a monopoly exists and their are no buildings. Otherwise, it's whatever the rent is
func (r *rentalProperty) CollectRent(rentee Landlord) error {
	if r.landlord != nil && !r.mortgaged {
		rent := r.rent.Rent(0)
		if r.monopoly {
			rent = rent * 2
		}
		if err:=rentee.Spend(rent,nil,nil); err !=nil{
			//rentee can't pay with cash. attempt to collect other assets
			cash := rentee.Cash();
		}
	}

}

//Mortgageable is no buildings, it is owned, and it's not already mortgaged
func (r *rentalProperty) Mortgagable() bool{
	return r.landlord != nil && !r.Mortgaged()
}

