package field

import (
	"bytes"
	"fmt"
	"math/rand"
	"sort"
	"strconv"
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
	Time      int

	rate float64
	mode int
}

func New(start int, end int, rate float64) *Field {
	rand.Seed(time.Now().UnixMicro())
	f := &Field{}

	f.start = start
	f.end = end
	f.Time = start
	f.rate = rate

	return f
}

func (f *Field) Loop() {
	for t := f.start; t < f.end; t++ {
		f.Time = t
		fmt.Printf("a\n")
		f.Step()
		time.Sleep(time.Second * 1)
	}
}

func (f *Field) Step() {
	for _, e := range f.Elevators {
		e.Step()
	}

	for _, h := range f.Humans {
		h.Step(f.Elevators)
	}

	f.NewHuman()
	f.Draw()
}

func (f *Field) NewHuman() {
	rand.Seed(time.Now().UnixNano())
	if rand.Float64() < f.rate {
		var cfl, tfl int
		cfl = rand.Intn(len(f.Floors))
		tfl = rand.Intn(len(f.Floors))
		for tfl == cfl {
			tfl = rand.Intn(len(f.Floors))
		}
		f.Humans = append(f.Humans, human.New(f.Floors[cfl], f.Floors[tfl]))
	}
}

func (f *Field) Draw() {
	f.Visualize()
}

func (f *Field) Visualize() {
	var keys []int
	strings := map[int]string{}

	for _, fl := range f.Floors {
		strings[fl.Number] = strconv.Itoa(fl.Number)
		strings[fl.Number] += " "
		if fl.Up {
			strings[fl.Number] += "↑"
		} else {
			strings[fl.Number] += " "
		}
		if fl.Down {
			strings[fl.Number] += "↓"
		} else {
			strings[fl.Number] += " "
		}
		strings[fl.Number] += " |"
		keys = append(keys, fl.Number)
	}

	sort.Ints(keys)

	for _, e := range f.Elevators {
		for _, k := range keys {
			if k == e.CurrentFloor.Number {
				if e.IsStoped {
					if e.Counter == e.Wait {
						strings[k] += "*"
					} else {
						strings[k] += strconv.Itoa(e.Wait - e.Counter)
					}
				} else if e.Up() {
					strings[k] += "↑"
				} else if e.Down() {
					strings[k] += "↓"
				} else {
					strings[k] += " "
				}
			} else {
				strings[k] += " "
			}
			strings[k] += "|"
		}
	}

	for i, j := 0, len(keys)-1; i < j; i, j = i+1, j-1 {
		keys[i], keys[j] = keys[j], keys[i]
	}

	for _, h := range f.Humans {
		if h.Arrival() {
			strings[h.CurrentFloor.Number] += "○"
		} else if h.Elevator == nil {
			strings[h.CurrentFloor.Number] += "*"
		}
	}

	var out bytes.Buffer

	out.WriteString(strconv.Itoa(f.Time))
	out.WriteString("\n")
	for _, k := range keys {
		out.WriteString(strings[k])
		out.WriteString("\n")
	}

	str := out.String()
	fmt.Printf(str)
}
