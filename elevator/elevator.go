package elevator

import (
	"github.com/kyuki3rain/elevator-simulator/floor"
)

type Elevator struct {
	capacity int
	speed    int
	Number   int

	People       int
	Hight        int
	CurrentFloor *floor.Floor
	targetFloors *[]floor.Floor
}

func New(currentFloor *floor.Floor, capacity int, number int) *Elevator {
	e := &Elevator{}

	e.CurrentFloor = currentFloor
	e.capacity = capacity
	e.speed = 2
	e.Number = number

	return e
}

func NewArray(currentFloors []*floor.Floor, capacities []int) []*Elevator {
	var elevators []*Elevator

	for i, cf := range currentFloors {
		elevators = append(elevators, New(cf, capacities[i], i+1))
	}

	return elevators
}

func (e *Elevator) IsStoped() bool {
	return len(*e.targetFloors) == 0
}

func (e *Elevator) IsSaturated() bool {
	return e.People == e.capacity
}
