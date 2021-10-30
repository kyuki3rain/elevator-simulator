package human

import (
	"bytes"
	"fmt"
	"math"
	"strconv"

	"github.com/kyuki3rain/elevator-simulator/elevator"
	"github.com/kyuki3rain/elevator-simulator/floor"
)

type Human struct {
	Number       int
	Elevator     *elevator.Elevator
	CurrentFloor *floor.Floor
	TargetFloor  *floor.Floor
}

func (h *Human) String() string {
	var out bytes.Buffer

	out.WriteString("人間")
	out.WriteString(strconv.Itoa(h.Number))
	out.WriteString("\n")
	if h.Elevator != nil {
		out.WriteString("エレベータ")
		out.WriteString(strconv.Itoa(h.Elevator.Number))
		out.WriteString("に乗車")
	} else {
		out.WriteString("エレベータに未乗車")
	}
	out.WriteString("\n")
	out.WriteString("現在：")
	out.WriteString(fmt.Sprintf("%s階", strconv.Itoa(h.CurrentFloor.Number)))
	out.WriteString("\n")
	out.WriteString("目標：")
	out.WriteString(fmt.Sprintf("%s階", strconv.Itoa(h.TargetFloor.Number)))
	out.WriteString("\n")
	out.WriteString("\n")

	return out.String()
}

func New(currentFloor *floor.Floor, targetFloor *floor.Floor, number int) *Human {

	h := &Human{}

	h.Number = number
	h.CurrentFloor = currentFloor
	h.TargetFloor = targetFloor

	return h
}

func NewArray(currentFloors []*floor.Floor, targetFloor []*floor.Floor) []*Human {
	var humans []*Human

	for i, cf := range currentFloors {
		humans = append(humans, New(cf, targetFloor[i], i+1))
	}

	return humans
}

func (h *Human) Push(es []*elevator.Elevator) {
	var nearElevator *elevator.Elevator
	switch h.CurrentFloor.Compare(h.TargetFloor) {
	case 1:
		if !h.CurrentFloor.Up {
			for i := range es {
				if es[i].Status == elevator.NotWorking || (es[i].Status == elevator.Up && es[i].Height < h.CurrentFloor.Height) {
					if nearElevator == nil || math.Abs(float64(nearElevator.Height-h.CurrentFloor.Height)) < math.Abs(float64(es[i].Height-h.CurrentFloor.Height)) {
						nearElevator = es[i]
					}
				}
			}

			if nearElevator != nil {
				h.CurrentFloor.Up = true
			}
		}
	case -1:
		if !h.CurrentFloor.Down {
			for i := range es {
				if es[i].Status == elevator.NotWorking || (es[i].Status == elevator.Down && es[i].Height > h.CurrentFloor.Height) {
					if nearElevator == nil || math.Abs(float64(nearElevator.Height-h.CurrentFloor.Height)) > math.Abs(float64(es[i].Height-h.CurrentFloor.Height)) {
						nearElevator = es[i]
					}
				}
			}

			if nearElevator != nil {
				h.CurrentFloor.Down = true
			}
		}
	}

	if nearElevator == nil {
		return
	}
	nearElevator.AddTargetFloor(h.CurrentFloor)
}

func (h *Human) Step(elevators []*elevator.Elevator) {
	if h.CurrentFloor == h.TargetFloor {
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

	status := h.CurrentFloor.Compare(h.TargetFloor)

	for _, e := range elevators {
		if e.CurrentFloor == h.CurrentFloor && e.IsStoped && !e.IsSaturated() && (e.Status == status || e.Status == elevator.NotWorking) {
			h.Elevator = e
			h.Elevator.AddTargetFloor(h.TargetFloor)
			h.Elevator.People++
			return
		}
	}

	h.Push(elevators)
}
