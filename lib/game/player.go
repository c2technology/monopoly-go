package game

import (
	"fmt"
)

type Player interface {
	fmt.Stringer
	Proprietor
	Builder
	Prisoner
	Buyer
}
