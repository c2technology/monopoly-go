package game

import (
	"fmt"
)

type Die interface {
	fmt.Stringer
	Roll()
	Value() int
}

type Dice interface {
	fmt.Stringer
	Roll()
	Doubles() bool
	Value() int
}
