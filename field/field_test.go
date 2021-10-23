package field

import (
	"testing"

	"github.com/kyuki3rain/elevator-simulator/elevator"
	"github.com/kyuki3rain/elevator-simulator/floor"
	"github.com/kyuki3rain/elevator-simulator/human"
)

func Test(t *testing.T) {
	f := New(0, 100)

	f.Floors = floor.NewArray([]int{1, 2, 3, 4, 5}, []int{5, 10, 15, 20, 25})
	f.Elevators = elevator.NewArray([]*floor.Floor{f.Floors[0], f.Floors[0], f.Floors[0]}, []int{1, 1, 1})
	f.Humans = human.NewArray([]*floor.Floor{f.Floors[0], f.Floors[0], f.Floors[1]}, []*floor.Floor{f.Floors[1], f.Floors[2], f.Floors[3]})

	tests := []int{1, 2, 0}

	f.Step()

	for i, h := range f.Humans {
		if h.Elevator != nil {
			if h.Elevator.Number != tests[i] {
				t.Errorf("h.Elevator is not matched.")
			}
			t.Log(h.Elevator.Number)
		} else {
			if tests[i] != 0 {
				t.Errorf("h.Elevator is not matched.")
			}
			t.Log("nil")
		}
	}
}
