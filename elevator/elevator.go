package elevator

import (
	"sort"

	"github.com/kyuki3rain/elevator-simulator/floor"
)

type Elevator struct {
	capacity int
	speed    int
	Number   int
	Wait     int

	People       int
	Height       int
	CurrentFloor *floor.Floor
	TargetFloors []*floor.Floor
	IsStoped     bool
	Counter      int
}

func New(currentFloor *floor.Floor, capacity int, number int) *Elevator {
	e := &Elevator{}

	e.CurrentFloor = currentFloor
	e.Height = e.CurrentFloor.Height
	e.capacity = capacity
	e.speed = 3
	e.Number = number
	e.Wait = 2
	e.IsStoped = true

	return e
}

func NewArray(currentFloors []*floor.Floor, capacities []int) []*Elevator {
	var elevators []*Elevator

	for i, cf := range currentFloors {
		elevators = append(elevators, New(cf, capacities[i], i+1))
	}

	return elevators
}

func (e *Elevator) Arrival(status int) bool {
	switch status {
	case 1:
		return e.TargetFloors[0].IsHigh(e.Height)
	case -1:
		return e.TargetFloors[0].IsLow(e.Height)
	}

	return true
}

func (e *Elevator) IsWorking() bool {
	return len(e.TargetFloors) != 0
}

func (e *Elevator) IsSaturated() bool {
	return e.People == e.capacity
}

func (e *Elevator) Up() bool {
	if e.IsStoped {
		return false
	}
	return e.CurrentFloor.Compare(e.TargetFloors[0]) == 1
}
func (e *Elevator) Down() bool {
	if e.IsStoped {
		return false
	}
	return e.CurrentFloor.Compare(e.TargetFloors[0]) == -1
}

func (e *Elevator) AddTargetFloor(f *floor.Floor) {
	if e.Up() {
		if e.CurrentFloor.Compare(f) == 1 {
			return
		}
	} else if e.Down() {
		if e.CurrentFloor.Compare(f) == -1 {
			return
		}
	}

	e.TargetFloors = append(e.TargetFloors, f)
	sort.Slice(e.TargetFloors, func(i, j int) bool { return e.TargetFloors[i].Compare(e.TargetFloors[j]) == 1 })
}

func (e *Elevator) Step() {
	if e.IsStoped {
		if e.IsWorking() {
			e.CurrentFloor.Arrive(e.CurrentFloor.Compare(e.TargetFloors[0]))
		}
		e.CurrentFloor.Arrive(0)
		if e.Counter >= e.Wait {
			if len(e.TargetFloors) == 0 {
				e.IsStoped = true
				return
			}
			e.Counter = 0
			e.IsStoped = false
		} else {
			e.Counter++
		}
		return
	}

	status := e.CurrentFloor.Compare(e.TargetFloors[0])

	switch status {
	case 1:
		e.Height += e.speed
		if e.CurrentFloor.UpFloor.Height <= e.Height {
			e.CurrentFloor = e.CurrentFloor.UpFloor
		}
	case -1:
		e.Height -= e.speed
		if e.CurrentFloor.DownFloor.Height >= e.Height {
			e.CurrentFloor = e.CurrentFloor.DownFloor
		}
	}

	if e.Arrival(status) {
		e.Height = e.CurrentFloor.Height
		e.IsStoped = true
		e.TargetFloors = e.TargetFloors[1:]
	}
}
