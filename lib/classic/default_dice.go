package classic

import (
	"fmt"
	"github.com/c2technology/monopoly-go/lib/game"
	"log"
	"math/rand"
)

func NewDice() game.Dice {
	dice := &dice{
		die1: newDie(6),
		die2: newDie(6),
	}
	log.Printf("Dice initialzied: %v", dice)
	return dice

}
func newDie(sides int) game.Die {
	if sides < 3 {
		sides = 3
	}
	values := []int{}
	for i := 0; i < sides; i++ {
		values = append(values, i+1)
	}
	current := values[rand.Intn(len(values))]
	return &die{
		values,
		current,
	}
}

type dice struct {
	die1      game.Die
	die2      game.Die
	doubles   bool
	faceValue int
}

func (d *dice) Roll() {
	d.die1.Roll()
	d.die2.Roll()
	d.doubles = d.die1.Value() == d.die2.Value()
	d.faceValue = d.die1.Value() + d.die2.Value()
}
func (d *dice) Value() int {
	return d.faceValue
}

func (d *dice) Doubles() bool {
	return d.doubles
}
func (d *dice) String() string {
	return fmt.Sprintf("%s + %s = %d (%t)", d.die1, d.die2, d.faceValue, d.doubles)
}

type die struct {
	values    []int
	faceValue int
}

func (d *die) String() string {
	return fmt.Sprintf("%d", d.faceValue)
}

func (d *die) Roll() {
	d.faceValue = d.values[rand.Intn(len(d.values))]
}

func (d *die) Value() int {
	return d.faceValue
}
