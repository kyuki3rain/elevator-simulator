package human

import (
	"github.com/kyuki3rain/elevator-simulator/elevator"
	"github.com/kyuki3rain/elevator-simulator/floor"
)

type Human struct {
	Elevator     *elevator.Elevator
	CurrentFloor *floor.Floor
	TargetFloor  *floor.Floor
}

func New(currentFloor *floor.Floor, targetFloor *floor.Floor) *Human {

	h := &Human{}

	h.CurrentFloor = currentFloor
	h.TargetFloor = targetFloor

	return h
}

func NewArray(currentFloors []*floor.Floor, targetFloor []*floor.Floor) []*Human {
	var humans []*Human

	for i, cf := range currentFloors {
		humans = append(humans, New(cf, targetFloor[i]))
	}

	return humans
}

func (h *Human) Ride(elevator *elevator.Elevator) {
	h.Elevator = elevator
	h.Elevator.People += 1
}

func (h *Human) Push() {
	if h.CurrentFloor.Compare(h.TargetFloor) {
		h.CurrentFloor.PushUp()
	} else {
		h.CurrentFloor.PushDown()
	}
}

func (h *Human) Step(elevators []*elevator.Elevator) {
	if h.Elevator != nil {
		return
	}

	for _, e := range elevators {
		if e.CurrentFloor == h.CurrentFloor && !e.IsSaturated() {
			h.Ride(e)
			return
		}
	}

	h.Push()
}
