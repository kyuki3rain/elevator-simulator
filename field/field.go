package field

import (
	"math/rand"
	"time"

	"github.com/kyuki3rain/elevator-simulator/elevator"
	"github.com/kyuki3rain/elevator-simulator/floor"
	"github.com/kyuki3rain/elevator-simulator/human"
)

type Field struct {
	Elevators []*elevator.Elevator
	Humans    []*human.Human
	Floors    []*floor.Floor
	start     int
	end       int
	time      int
}

func New(start int, end int) *Field {
	rand.Seed(time.Now().UnixMicro())
	f := &Field{}

	f.start = start
	f.end = end
	f.time = start

	return f
}

func (f *Field) Loop() {

	for time := f.start; time < f.end; time++ {
		f.time = time
		f.Step()
	}
}

func (f *Field) Step() {
	for _, h := range f.Humans {
		h.Step(f.Elevators)
	}
}
