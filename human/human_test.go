package human

import (
	"testing"

	"github.com/kyuki3rain/elevator-simulator/elevator"
	"github.com/kyuki3rain/elevator-simulator/floor"
)

func TestStep(t *testing.T) {
	floors := floor.NewArray([]int{1, 2, 3, 4, 5}, []int{5, 10, 15, 20, 25})
	elevators := elevator.NewArray([]*floor.Floor{floors[0], floors[0], floors[0]}, []int{1, 1, 1})
	humans := NewArray([]*floor.Floor{floors[0], floors[0], floors[1]}, []*floor.Floor{floors[1], floors[2], floors[3]})

	tests := []int{1, 2, 0}

	for i, h := range humans {
		h.Step(elevators)

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

	ftestsup := []bool{false, true, false, false, false}
	ftestsdown := []bool{false, false, false, false, false}

	for i, f := range floors {
		if f.Up() == ftestsup[i] && f.Down() == ftestsdown[i] {
			continue
		}

		t.Errorf("floor not match.")
	}
}
