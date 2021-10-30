package elevator

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"

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
	Counter      int
	Status       int
	IsStoped     bool
}

const (
	Up         = 1
	NotWorking = 0
	Down       = -1
)

func (e *Elevator) String() string {
	var out bytes.Buffer

	out.WriteString(fmt.Sprintf("%d番エレベータ\n", e.Number))
	out.WriteString("人数：")
	out.WriteString(strconv.Itoa(e.People))
	out.WriteString("\n")
	out.WriteString("高さ：")
	out.WriteString(strconv.Itoa(e.Height))
	out.WriteString("\n")
	out.WriteString("現在地：")
	out.WriteString(fmt.Sprintf("%s階", strconv.Itoa(e.CurrentFloor.Number)))
	out.WriteString("\n")
	out.WriteString("停止予定：")
	for i := range e.TargetFloors {
		if i != 0 {
			out.WriteString(", ")
		}
		out.WriteString(fmt.Sprintf("%s階", strconv.Itoa(e.TargetFloors[i].Number)))
	}
	out.WriteString("\n")
	out.WriteString("状態：")
	switch e.Status {
	case Up:
		out.WriteString("上昇")
	case NotWorking:
		out.WriteString("無し")
	case Down:
		out.WriteString("下降")
	}
	if e.IsStoped {
		out.WriteString(", 停車中")
	}
	out.WriteString("\n")
	out.WriteString("\n")

	return out.String()
}

func New(currentFloor *floor.Floor, capacity int, number int) *Elevator {
	e := &Elevator{}

	e.CurrentFloor = currentFloor
	e.Height = e.CurrentFloor.Height
	e.capacity = capacity
	e.speed = 3
	e.Number = number
	e.Wait = 2
	e.Status = NotWorking
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

func (e *Elevator) Arrival() bool {
	switch e.Status {
	case Up:
		return e.TargetFloors[0].Height <= e.Height
	case NotWorking:
		return true
	case Down:
		return e.TargetFloors[len(e.TargetFloors)-1].Height >= e.Height
	}

	return true
}

func (e *Elevator) NoTarget() bool {
	return len(e.TargetFloors) == 0
}

func (e *Elevator) IsSaturated() bool {
	return e.People == e.capacity
}

func (e *Elevator) AddTargetFloor(f *floor.Floor) {
	status := e.CurrentFloor.Compare(f)

	if status == 0 {
		return
	}

	switch e.Status {
	case Up:
		if status == Down {
			panic("add down target to up elevator")
		}
	case NotWorking:
		e.Status = status
	case Down:
		if status == Up {
			panic("add up target to down elevator")
		}
	}

	e.TargetFloors = append(e.TargetFloors, f)
	e.TargetFloors = removeDuplicate1(e.TargetFloors)
	sort.Slice(e.TargetFloors, func(i, j int) bool { return e.TargetFloors[i].Compare(e.TargetFloors[j]) == 1 })
}

func removeDuplicate1(args []*floor.Floor) []*floor.Floor {
	results := make([]*floor.Floor, 0, len(args))
	encountered := map[int]bool{}
	for i := 0; i < len(args); i++ {
		if !encountered[args[i].Number] {
			encountered[args[i].Number] = true
			results = append(results, args[i])
		}
	}
	return results
}

func (e *Elevator) Step() {
	if e.IsStoped {
		e.CurrentFloor.Arrive(e.Status)
		if e.Counter >= e.Wait {
			if e.Status == NotWorking {
				return
			}

			e.Counter = 0
			e.IsStoped = false
		} else {
			e.Counter++
		}
		return
	}

	switch e.Status {
	case Up:
		e.Height += e.speed
		for e.CurrentFloor.UpFloor != nil && e.CurrentFloor.UpFloor.Height <= e.Height {
			e.CurrentFloor = e.CurrentFloor.UpFloor
		}
		if e.Arrival() {
			e.Height = e.CurrentFloor.Height
			for !e.NoTarget() && e.TargetFloors[0] == e.CurrentFloor {
				e.TargetFloors = e.TargetFloors[1:]
			}
			e.IsStoped = true
		}
	case Down:
		e.Height -= e.speed
		for e.CurrentFloor.DownFloor != nil && e.CurrentFloor.DownFloor.Height >= e.Height {
			e.CurrentFloor = e.CurrentFloor.DownFloor
		}
		if e.Arrival() {
			e.Height = e.CurrentFloor.Height
			for !e.NoTarget() && e.TargetFloors[len(e.TargetFloors)-1] == e.CurrentFloor {
				if len(e.TargetFloors) == 1 {
					e.TargetFloors = []*floor.Floor{}
				} else {
					e.TargetFloors = e.TargetFloors[:len(e.TargetFloors)-2]
				}
			}
			e.IsStoped = true
		}
	}

	if e.NoTarget() {
		e.Status = NotWorking
	}
}
