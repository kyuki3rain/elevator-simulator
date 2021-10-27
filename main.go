package main

import (
	"github.com/kyuki3rain/elevator-simulator/elevator"
	"github.com/kyuki3rain/elevator-simulator/field"
	"github.com/kyuki3rain/elevator-simulator/floor"
	"github.com/kyuki3rain/elevator-simulator/human"
)

func main() {
	f := field.New(1, 100, 0.2)

	f.Floors = floor.NewArray([]int{1, 2, 3, 4}, []int{0, 5, 10, 15})
	f.Elevators = elevator.NewArray([]*floor.Floor{f.Floors[0], f.Floors[0], f.Floors[0]}, []int{2, 2, 2})
	f.Humans = human.NewArray([]*floor.Floor{f.Floors[0], f.Floors[0], f.Floors[1]}, []*floor.Floor{f.Floors[2], f.Floors[1], f.Floors[3]})

	f.Time = 0
	f.Draw()
	f.Loop()
}
