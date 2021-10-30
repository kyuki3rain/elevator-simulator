package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/kyuki3rain/elevator-simulator/elevator"
	"github.com/kyuki3rain/elevator-simulator/field"
	"github.com/kyuki3rain/elevator-simulator/floor"
	"github.com/kyuki3rain/elevator-simulator/human"
)

func main() {
	f := field.New(1, 1000, 10)

	f.Floors = floor.NewArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{0, 5, 10, 15, 20, 25, 30, 35, 40})
	f.Elevators = elevator.NewArray([]*floor.Floor{f.Floors[0], f.Floors[0], f.Floors[0], f.Floors[0], f.Floors[0]}, []int{100, 100, 100, 100, 100})
	f.Humans = human.NewArray([]*floor.Floor{f.Floors[0]}, []*floor.Floor{f.Floors[1]})

	f.Time = 0
	f.Draw()

	go func() {
		f.Loop(10)
		os.Exit(0)
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Println(f.String())
}
