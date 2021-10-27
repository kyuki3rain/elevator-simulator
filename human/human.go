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
	h.Elevator.People++
	h.Elevator.AddTargetFloor(h.TargetFloor)
}

func (h *Human) Push(es []*elevator.Elevator) {
	switch h.CurrentFloor.Compare(h.TargetFloor) {
	case 1:
		if !h.CurrentFloor.Up {
			var nearElevator *elevator.Elevator
			h.CurrentFloor.Up = true
			for i := range es {
				if !es[i].IsWorking() || (es[i].Up() && h.CurrentFloor.Compare(es[i].TargetFloors[len(es[i].TargetFloors)-1]) != -1) {
					if nearElevator == nil || nearElevator.Height < es[i].Height {
						nearElevator = es[i]
					}
				}
			}

			nearElevator.AddTargetFloor(h.CurrentFloor)
		}
		return
	case 0:
		return
	case -1:
		if !h.CurrentFloor.Down {
			var nearElevator *elevator.Elevator
			h.CurrentFloor.Down = true
			for i := range es {
				if !es[i].IsWorking() || (es[i].Down() && h.CurrentFloor.Compare(es[i].TargetFloors[len(es[i].TargetFloors)-1]) != 1) {
					if nearElevator == nil || nearElevator.Height > es[i].Height {
						nearElevator = es[i]
					}
				}
			}

			nearElevator.AddTargetFloor(h.CurrentFloor)
		}
		return
	}
}

func (h *Human) Arrival() bool {
	return h.CurrentFloor == h.TargetFloor
}

func (h *Human) Step(elevators []*elevator.Elevator) {
	if h.Arrival() {
		if h.Elevator != nil {
			h.Elevator.People--
			h.Elevator = nil
		}
		return
	}

	if h.Elevator != nil {
		h.CurrentFloor = h.Elevator.CurrentFloor
		return
	}

	for _, e := range elevators {
		if e.CurrentFloor == h.CurrentFloor && !e.IsSaturated() {
			h.Ride(e)
			return
		}
	}

	h.Push(elevators)
}
