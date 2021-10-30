package field

import (
	"testing"

	"github.com/kyuki3rain/elevator-simulator/elevator"
	"github.com/kyuki3rain/elevator-simulator/floor"
	"github.com/kyuki3rain/elevator-simulator/human"
)

func TestUp(t *testing.T) {
	f := New(1, 20, 0, 2, 1, 1, 0)

	f.Floors = floor.NewArray([]int{1, 2}, []int{0, 5})
	f.Elevators = elevator.NewArray([]*floor.Floor{f.Floors[0]}, []int{1})
	f.Humans = human.NewArray([]*floor.Floor{f.Floors[0]}, []*floor.Floor{f.Floors[1]})

	f.Draw()
	if !testElevator(t, f.Elevators[0], f.Floors[0], 0, 0, []*floor.Floor{}, true, 0) || !testHuman(t, f.Humans[0], nil, f.Floors[0]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], false, false) {
		return
	}

	f.Time = 1
	f.Step()
	if !testElevator(t, f.Elevators[0], f.Floors[0], 1, 0, []*floor.Floor{f.Floors[1]}, true, 1) || !testHuman(t, f.Humans[0], f.Elevators[0], f.Floors[0]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], false, false) {
		return
	}

	f.Time = 2
	f.Step()
	if !testElevator(t, f.Elevators[0], f.Floors[0], 1, 0, []*floor.Floor{f.Floors[1]}, true, 2) || !testHuman(t, f.Humans[0], f.Elevators[0], f.Floors[0]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], false, false) {
		return
	}

	f.Time = 3
	f.Step()
	if !testElevator(t, f.Elevators[0], f.Floors[0], 1, 0, []*floor.Floor{f.Floors[1]}, false, 0) || !testHuman(t, f.Humans[0], f.Elevators[0], f.Floors[0]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], false, false) {
		return
	}

	f.Time = 4
	f.Step()
	if !testElevator(t, f.Elevators[0], f.Floors[0], 1, 3, []*floor.Floor{f.Floors[1]}, false, 0) || !testHuman(t, f.Humans[0], f.Elevators[0], f.Floors[0]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], false, false) {
		return
	}

	f.Time = 5
	f.Step()
	if !testElevator(t, f.Elevators[0], f.Floors[1], 1, 5, []*floor.Floor{}, true, 0) || !testHuman(t, f.Humans[0], f.Elevators[0], f.Floors[1]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], false, false) {
		return
	}

	f.Time = 6
	f.Step()
	if !testElevator(t, f.Elevators[0], f.Floors[1], 0, 5, []*floor.Floor{}, true, 1) || !testHuman(t, f.Humans[0], nil, f.Floors[1]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], false, false) {
		return
	}

	f.Time = 7
	f.Step()
	if !testElevator(t, f.Elevators[0], f.Floors[1], 0, 5, []*floor.Floor{}, true, 2) || !testHuman(t, f.Humans[0], nil, f.Floors[1]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], false, false) {
		return
	}

	f.Time = 8
	f.Step()
	if !testElevator(t, f.Elevators[0], f.Floors[1], 0, 5, []*floor.Floor{}, true, 2) || !testHuman(t, f.Humans[0], nil, f.Floors[1]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], false, false) {
		return
	}
}

func TestDown(t *testing.T) {
	f := New(1, 20, 0, 2, 1, 1, 0)

	f.Floors = floor.NewArray([]int{1, 2}, []int{0, 5})
	f.Elevators = elevator.NewArray([]*floor.Floor{f.Floors[1]}, []int{1})
	f.Humans = human.NewArray([]*floor.Floor{f.Floors[1]}, []*floor.Floor{f.Floors[0]})

	f.Draw()
	if !testElevator(t, f.Elevators[0], f.Floors[1], 0, 5, []*floor.Floor{}, true, 0) || !testHuman(t, f.Humans[0], nil, f.Floors[1]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], false, false) {
		return
	}

	f.Time = 1
	f.Step()
	if !testElevator(t, f.Elevators[0], f.Floors[1], 1, 5, []*floor.Floor{f.Floors[0]}, true, 1) || !testHuman(t, f.Humans[0], f.Elevators[0], f.Floors[1]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], false, false) {
		return
	}

	f.Time = 2
	f.Step()
	if !testElevator(t, f.Elevators[0], f.Floors[1], 1, 5, []*floor.Floor{f.Floors[0]}, true, 2) || !testHuman(t, f.Humans[0], f.Elevators[0], f.Floors[1]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], false, false) {
		return
	}

	f.Time = 3
	f.Step()
	if !testElevator(t, f.Elevators[0], f.Floors[1], 1, 5, []*floor.Floor{f.Floors[0]}, false, 0) || !testHuman(t, f.Humans[0], f.Elevators[0], f.Floors[1]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], false, false) {
		return
	}

	f.Time = 4
	f.Step()
	if !testElevator(t, f.Elevators[0], f.Floors[1], 1, 2, []*floor.Floor{f.Floors[0]}, false, 0) || !testHuman(t, f.Humans[0], f.Elevators[0], f.Floors[1]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], false, false) {
		return
	}

	f.Time = 5
	f.Step()
	if !testElevator(t, f.Elevators[0], f.Floors[0], 1, 0, []*floor.Floor{}, true, 0) || !testHuman(t, f.Humans[0], f.Elevators[0], f.Floors[0]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], false, false) {
		return
	}

	f.Time = 6
	f.Step()
	if !testElevator(t, f.Elevators[0], f.Floors[0], 0, 0, []*floor.Floor{}, true, 1) || !testHuman(t, f.Humans[0], nil, f.Floors[0]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], false, false) {
		return
	}

	f.Time = 7
	f.Step()
	if !testElevator(t, f.Elevators[0], f.Floors[0], 0, 0, []*floor.Floor{}, true, 2) || !testHuman(t, f.Humans[0], nil, f.Floors[0]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], false, false) {
		return
	}

	f.Time = 8
	f.Step()
	if !testElevator(t, f.Elevators[0], f.Floors[0], 0, 0, []*floor.Floor{}, true, 2) || !testHuman(t, f.Humans[0], nil, f.Floors[0]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], false, false) {
		return
	}
}

func TestUpButton(t *testing.T) {
	f := New(1, 20, 0, 2, 1, 1, 0)

	f.Floors = floor.NewArray([]int{1, 2}, []int{0, 5})
	f.Elevators = elevator.NewArray([]*floor.Floor{f.Floors[1]}, []int{1})
	f.Humans = human.NewArray([]*floor.Floor{f.Floors[0]}, []*floor.Floor{f.Floors[1]})

	f.Draw()
	if !testElevator(t, f.Elevators[0], f.Floors[1], 0, 5, []*floor.Floor{}, true, 0) || !testHuman(t, f.Humans[0], nil, f.Floors[0]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], false, false) {
		return
	}

	f.Time = 1
	f.Step()
	if !testElevator(t, f.Elevators[0], f.Floors[1], 0, 5, []*floor.Floor{f.Floors[0]}, true, 1) || !testHuman(t, f.Humans[0], nil, f.Floors[0]) || !testFloor(t, f.Floors[0], true, false) || !testFloor(t, f.Floors[1], false, false) {
		return
	}

	f.Time = 2
	f.Step()
	if !testElevator(t, f.Elevators[0], f.Floors[1], 0, 5, []*floor.Floor{f.Floors[0]}, true, 2) || !testHuman(t, f.Humans[0], nil, f.Floors[0]) || !testFloor(t, f.Floors[0], true, false) || !testFloor(t, f.Floors[1], false, false) {
		return
	}

	f.Time = 3
	f.Step()
	if !testElevator(t, f.Elevators[0], f.Floors[1], 0, 5, []*floor.Floor{f.Floors[0]}, false, 0) || !testHuman(t, f.Humans[0], nil, f.Floors[0]) || !testFloor(t, f.Floors[0], true, false) || !testFloor(t, f.Floors[1], false, false) {
		return
	}

	f.Time = 4
	f.Step()
	if !testElevator(t, f.Elevators[0], f.Floors[1], 0, 2, []*floor.Floor{f.Floors[0]}, false, 0) || !testHuman(t, f.Humans[0], nil, f.Floors[0]) || !testFloor(t, f.Floors[0], true, false) || !testFloor(t, f.Floors[1], false, false) {
		return
	}

	f.Time = 5
	f.Step()
	if !testElevator(t, f.Elevators[0], f.Floors[0], 1, 0, []*floor.Floor{f.Floors[1]}, true, 0) || !testHuman(t, f.Humans[0], f.Elevators[0], f.Floors[0]) || !testFloor(t, f.Floors[0], true, false) || !testFloor(t, f.Floors[1], false, false) {
		return
	}

	f.Time = 6
	f.Step()
	if !testElevator(t, f.Elevators[0], f.Floors[0], 1, 0, []*floor.Floor{f.Floors[1]}, true, 1) || !testHuman(t, f.Humans[0], f.Elevators[0], f.Floors[0]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], false, false) {
		return
	}

	f.Time = 7
	f.Step()
	if !testElevator(t, f.Elevators[0], f.Floors[0], 1, 0, []*floor.Floor{f.Floors[1]}, true, 2) || !testHuman(t, f.Humans[0], f.Elevators[0], f.Floors[0]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], false, false) {
		return
	}

	f.Time = 8
	f.Step()
	if !testElevator(t, f.Elevators[0], f.Floors[0], 1, 0, []*floor.Floor{f.Floors[1]}, false, 0) || !testHuman(t, f.Humans[0], f.Elevators[0], f.Floors[0]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], false, false) {
		return
	}

	f.Time = 9
	f.Step()
	if !testElevator(t, f.Elevators[0], f.Floors[0], 1, 3, []*floor.Floor{f.Floors[1]}, false, 0) || !testHuman(t, f.Humans[0], f.Elevators[0], f.Floors[0]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], false, false) {
		return
	}

	f.Time = 10
	f.Step()
	if !testElevator(t, f.Elevators[0], f.Floors[1], 1, 5, []*floor.Floor{}, true, 0) || !testHuman(t, f.Humans[0], f.Elevators[0], f.Floors[1]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], false, false) {
		return
	}

	f.Time = 11
	f.Step()
	if !testElevator(t, f.Elevators[0], f.Floors[1], 0, 5, []*floor.Floor{}, true, 1) || !testHuman(t, f.Humans[0], nil, f.Floors[1]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], false, false) {
		return
	}
}

func TestDownButton(t *testing.T) {
	f := New(1, 20, 0, 2, 1, 1, 0)

	f.Floors = floor.NewArray([]int{1, 2}, []int{0, 5})
	f.Elevators = elevator.NewArray([]*floor.Floor{f.Floors[0]}, []int{1})
	f.Humans = human.NewArray([]*floor.Floor{f.Floors[1]}, []*floor.Floor{f.Floors[0]})

	f.Draw()
	if !testElevator(t, f.Elevators[0], f.Floors[0], 0, 0, []*floor.Floor{}, true, 0) || !testHuman(t, f.Humans[0], nil, f.Floors[1]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], false, false) {
		return
	}

	f.Time = 1
	f.Step()
	if !testElevator(t, f.Elevators[0], f.Floors[0], 0, 0, []*floor.Floor{f.Floors[1]}, true, 1) || !testHuman(t, f.Humans[0], nil, f.Floors[1]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], false, true) {
		return
	}

	f.Time = 2
	f.Step()
	if !testElevator(t, f.Elevators[0], f.Floors[0], 0, 0, []*floor.Floor{f.Floors[1]}, true, 2) || !testHuman(t, f.Humans[0], nil, f.Floors[1]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], false, true) {
		return
	}

	f.Time = 3
	f.Step()
	if !testElevator(t, f.Elevators[0], f.Floors[0], 0, 0, []*floor.Floor{f.Floors[1]}, false, 0) || !testHuman(t, f.Humans[0], nil, f.Floors[1]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], false, true) {
		return
	}

	f.Time = 4
	f.Step()
	if !testElevator(t, f.Elevators[0], f.Floors[0], 0, 3, []*floor.Floor{f.Floors[1]}, false, 0) || !testHuman(t, f.Humans[0], nil, f.Floors[1]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], false, true) {
		return
	}

	f.Time = 5
	f.Step()
	if !testElevator(t, f.Elevators[0], f.Floors[1], 1, 5, []*floor.Floor{f.Floors[0]}, true, 0) || !testHuman(t, f.Humans[0], f.Elevators[0], f.Floors[1]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], false, true) {
		return
	}

	f.Time = 6
	f.Step()
	if !testElevator(t, f.Elevators[0], f.Floors[1], 1, 5, []*floor.Floor{f.Floors[0]}, true, 1) || !testHuman(t, f.Humans[0], f.Elevators[0], f.Floors[1]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], false, false) {
		return
	}

	f.Time = 7
	f.Step()
	if !testElevator(t, f.Elevators[0], f.Floors[1], 1, 5, []*floor.Floor{f.Floors[0]}, true, 2) || !testHuman(t, f.Humans[0], f.Elevators[0], f.Floors[1]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], false, false) {
		return
	}

	f.Time = 8
	f.Step()
	if !testElevator(t, f.Elevators[0], f.Floors[1], 1, 5, []*floor.Floor{f.Floors[0]}, false, 0) || !testHuman(t, f.Humans[0], f.Elevators[0], f.Floors[1]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], false, false) {
		return
	}

	f.Time = 9
	f.Step()
	if !testElevator(t, f.Elevators[0], f.Floors[1], 1, 2, []*floor.Floor{f.Floors[0]}, false, 0) || !testHuman(t, f.Humans[0], f.Elevators[0], f.Floors[1]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], false, false) {
		return
	}

	f.Time = 10
	f.Step()
	if !testElevator(t, f.Elevators[0], f.Floors[0], 1, 0, []*floor.Floor{}, true, 0) || !testHuman(t, f.Humans[0], f.Elevators[0], f.Floors[0]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], false, false) {
		return
	}

	f.Time = 11
	f.Step()
	if !testElevator(t, f.Elevators[0], f.Floors[0], 0, 0, []*floor.Floor{}, true, 1) || !testHuman(t, f.Humans[0], nil, f.Floors[0]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], false, false) {
		return
	}
}

func TestTwoElevator(t *testing.T) {
	f := New(1, 20, 0, 2, 1, 1, 0)

	f.Floors = floor.NewArray([]int{1, 2}, []int{0, 5})
	f.Elevators = elevator.NewArray([]*floor.Floor{f.Floors[0], f.Floors[0]}, []int{1, 1})
	f.Humans = human.NewArray([]*floor.Floor{f.Floors[1]}, []*floor.Floor{f.Floors[0]})

	f.Draw()
	if !testElevator(t, f.Elevators[0], f.Floors[0], 0, 0, []*floor.Floor{}, true, 0) || !testHuman(t, f.Humans[0], nil, f.Floors[1]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], false, false) {
		return
	}

	f.Time = 1
	f.Step()
	if !testElevator(t, f.Elevators[0], f.Floors[0], 0, 0, []*floor.Floor{f.Floors[1]}, true, 1) || !testElevator(t, f.Elevators[1], f.Floors[0], 0, 0, []*floor.Floor{}, true, 1) || !testHuman(t, f.Humans[0], nil, f.Floors[1]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], false, true) {
		return
	}

	f.Time = 2
	f.Step()
	if !testElevator(t, f.Elevators[0], f.Floors[0], 0, 0, []*floor.Floor{f.Floors[1]}, true, 2) || !testElevator(t, f.Elevators[1], f.Floors[0], 0, 0, []*floor.Floor{}, true, 2) || !testHuman(t, f.Humans[0], nil, f.Floors[1]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], false, true) {
		return
	}

	f.Time = 3
	f.Step()
	if !testElevator(t, f.Elevators[0], f.Floors[0], 0, 0, []*floor.Floor{f.Floors[1]}, false, 0) || !testElevator(t, f.Elevators[1], f.Floors[0], 0, 0, []*floor.Floor{}, true, 2) || !testHuman(t, f.Humans[0], nil, f.Floors[1]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], false, true) {
		return
	}
}

func TestTwoHuman(t *testing.T) {
	f := New(1, 20, 0, 2, 1, 1, 0)

	f.Floors = floor.NewArray([]int{1, 2, 3}, []int{0, 5, 10})
	f.Elevators = elevator.NewArray([]*floor.Floor{f.Floors[0]}, []int{1})
	f.Humans = human.NewArray([]*floor.Floor{f.Floors[0], f.Floors[1]}, []*floor.Floor{f.Floors[1], f.Floors[2]})

	f.Draw()
	if !testElevator(t, f.Elevators[0], f.Floors[0], 0, 0, []*floor.Floor{}, true, 0) || !testHuman(t, f.Humans[0], nil, f.Floors[0]) || !testHuman(t, f.Humans[1], nil, f.Floors[1]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], false, false) || !testFloor(t, f.Floors[2], false, false) {
		return
	}

	f.Time = 1
	f.Step()
	if !testElevator(t, f.Elevators[0], f.Floors[0], 1, 0, []*floor.Floor{f.Floors[1]}, true, 1) || !testHuman(t, f.Humans[0], f.Elevators[0], f.Floors[0]) || !testHuman(t, f.Humans[1], nil, f.Floors[1]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], true, false) || !testFloor(t, f.Floors[2], false, false) {
		return
	}

	f.Time = 2
	f.Step()
	if !testElevator(t, f.Elevators[0], f.Floors[0], 1, 0, []*floor.Floor{f.Floors[1]}, true, 2) || !testHuman(t, f.Humans[0], f.Elevators[0], f.Floors[0]) || !testHuman(t, f.Humans[1], nil, f.Floors[1]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], true, false) || !testFloor(t, f.Floors[2], false, false) {
		return
	}

	f.Time = 3
	f.Step()
	if !testElevator(t, f.Elevators[0], f.Floors[0], 1, 0, []*floor.Floor{f.Floors[1]}, false, 0) || !testHuman(t, f.Humans[0], f.Elevators[0], f.Floors[0]) || !testHuman(t, f.Humans[1], nil, f.Floors[1]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], true, false) || !testFloor(t, f.Floors[2], false, false) {
		return
	}

	f.Time = 4
	f.Step()
	if !testElevator(t, f.Elevators[0], f.Floors[0], 1, 3, []*floor.Floor{f.Floors[1]}, false, 0) || !testHuman(t, f.Humans[0], f.Elevators[0], f.Floors[0]) || !testHuman(t, f.Humans[1], nil, f.Floors[1]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], true, false) || !testFloor(t, f.Floors[2], false, false) {
		return
	}

	f.Time = 5
	f.Step()
	if !testElevator(t, f.Elevators[0], f.Floors[1], 1, 5, []*floor.Floor{}, true, 0) || !testHuman(t, f.Humans[0], f.Elevators[0], f.Floors[1]) || !testHuman(t, f.Humans[1], nil, f.Floors[1]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], true, false) || !testFloor(t, f.Floors[2], false, false) {
		return
	}

	f.Time = 6
	f.Step()
	if !testElevator(t, f.Elevators[0], f.Floors[1], 1, 5, []*floor.Floor{f.Floors[2]}, true, 1) || !testHuman(t, f.Humans[0], nil, f.Floors[1]) || !testHuman(t, f.Humans[1], f.Elevators[0], f.Floors[1]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], true, false) || !testFloor(t, f.Floors[2], false, false) {
		return
	}

	f.Time = 7
	f.Step()
	if !testElevator(t, f.Elevators[0], f.Floors[1], 1, 5, []*floor.Floor{f.Floors[2]}, true, 2) || !testHuman(t, f.Humans[0], nil, f.Floors[1]) || !testHuman(t, f.Humans[1], f.Elevators[0], f.Floors[1]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], false, false) || !testFloor(t, f.Floors[2], false, false) {
		return
	}

	f.Time = 8
	f.Step()
	if !testElevator(t, f.Elevators[0], f.Floors[1], 1, 5, []*floor.Floor{f.Floors[2]}, false, 0) || !testHuman(t, f.Humans[0], nil, f.Floors[1]) || !testHuman(t, f.Humans[1], f.Elevators[0], f.Floors[1]) || !testFloor(t, f.Floors[0], false, false) || !testFloor(t, f.Floors[1], false, false) || !testFloor(t, f.Floors[2], false, false) {
		return
	}
}

func testHuman(t *testing.T, Human *human.Human, Elevator *elevator.Elevator, CurrentFloor *floor.Floor) bool {
	r := true

	if Human.Elevator != Elevator {
		if Human.Elevator != nil {
			if Elevator != nil {
				t.Errorf("Elevator expect=%d, got=%d\n", Elevator.Number, Human.Elevator.Number)
			} else {
				t.Errorf("Elevator expect=nil, got=%d\n", Human.Elevator.Number)
			}
		} else if Elevator != nil {
			t.Errorf("Elevator expect=%d, got=nil\n", Elevator.Number)
		} else {
			t.Errorf("Elevator expect=nil, got=nil\n")
		}
		r = false
	}
	if Human.CurrentFloor != CurrentFloor {
		t.Errorf("CurrentFloor expect=%d, got=%d\n", CurrentFloor.Number, Human.CurrentFloor.Number)
		r = false
	}

	return r
}

func testElevator(t *testing.T, Elevator *elevator.Elevator, CurrentFloor *floor.Floor, People int, Height int, TargetFloors []*floor.Floor, IsStoped bool, Counter int) bool {
	r := true

	if Elevator.CurrentFloor != CurrentFloor {
		t.Errorf("CurrentFloor expect=%d, got=%d\n", CurrentFloor.Number, Elevator.CurrentFloor.Number)
		r = false
	}
	if Elevator.People != People {
		t.Errorf("People expect=%d, got=%d\n", People, Elevator.People)
		r = false
	}
	if Elevator.Height != Height {
		t.Errorf("Height expect=%d, got=%d\n", Height, Elevator.Height)
		r = false
	}

	if len(Elevator.TargetFloors) != len(TargetFloors) {
		t.Errorf("len(TargetFloors) expect=%d, got=%d\n", len(TargetFloors), len(Elevator.TargetFloors))
		r = false
	}

	for i := range Elevator.TargetFloors {
		if Elevator.TargetFloors[i] != TargetFloors[i] {
			t.Errorf("TargetFloors[%d] expect=%d, got=%d\n", i, TargetFloors[i].Number, Elevator.TargetFloors[i].Number)
			r = false
		}
	}
	if Elevator.IsStoped != IsStoped {
		t.Errorf("IsStoped expect=%t, got=%t\n", IsStoped, Elevator.IsStoped)
		r = false
	}
	if Elevator.Counter != Counter {
		t.Errorf("Counter expect=%d, got=%d\n", Counter, Elevator.Counter)
		r = false
	}

	return r
}

func testFloor(t *testing.T, fl *floor.Floor, up bool, down bool) bool {
	r := true

	if fl.Up != up {
		t.Errorf("Up expect=%t, got=%t\n", up, fl.Up)
		r = false
	}

	if fl.Down != down {
		t.Errorf("Up expect=%t, got=%t\n", down, fl.Down)
		r = false
	}

	return r
}
