package field

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/signal"
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

	freq float64
}

func New(start int, end int, rate float64, fl int, el int, max int, sleep int) *Field {
	rand.Seed(time.Now().UnixMicro())
	f := &Field{}

	f.start = start
	f.end = end
	f.Time = start
	f.freq = rate

	f.Floors = floor.NewArray(CreateRangeSlice(1, fl, 1), CreateRangeSlice(0, (fl-1)*5, 5))
	f.Elevators = elevator.NewArray(CreateFloorSlice(el, f.Floors[0]), CreateNumberSlice(el, max))
	f.Humans = human.NewArray([]*floor.Floor{f.Floors[0]}, []*floor.Floor{f.Floors[fl-1]})

	return f
}

func CreateFloorSlice(num int, fl *floor.Floor) []*floor.Floor {
	var res []*floor.Floor
	for i := 0; i < num; i++ {
		res = append(res, fl)
	}

	return res
}

func CreateNumberSlice(num int, t int) []int {
	var res []int
	for i := 0; i < num; i++ {
		res = append(res, t)
	}

	return res
}

func CreateRangeSlice(start int, end int, step int) []int {
	var res []int
	for i := start; i <= end; i += step {
		res = append(res, i)
	}

	return res
}

func (f *Field) Loop(sleep int) {
	for t := f.start; t < f.end; t++ {
		f.Time = t
		f.Step()
		time.Sleep(time.Millisecond * time.Duration(sleep))
	}
}

func (f *Field) LoopWithTest(sleep int) {
	go func() {
		f.Loop(sleep)
		os.Exit(0)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Println(f.String())
}

func (f *Field) Step() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Recover!:", err)
			fmt.Println(f.String())
		}
	}()

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
	for i := 0; i < int(math.Abs(rand.NormFloat64())*f.freq); i++ {
		var cfl, tfl int
		cfl = rand.Intn(len(f.Floors))
		tfl = rand.Intn(len(f.Floors))
		for tfl == cfl {
			tfl = rand.Intn(len(f.Floors))
		}
		f.Humans = append(f.Humans, human.New(f.Floors[cfl], f.Floors[tfl], f.Humans[len(f.Humans)-1].Number+1))
	}
}

func (f *Field) String() string {
	var out bytes.Buffer
	for _, fl := range f.Floors {
		out.WriteString(fl.String())
	}
	for _, e := range f.Elevators {
		out.WriteString(e.String())
	}
	for _, h := range f.Humans {
		out.WriteString(h.String())
	}

	return out.String()
}

func (f *Field) Draw() {
	var keys []int
	strings := map[int]string{}
	elevInfo := ""

	for _, fl := range f.Floors {
		if fl.Number < 10 {
			strings[fl.Number] = " "
		} else {
			strings[fl.Number] = ""
		}
		strings[fl.Number] += strconv.Itoa(fl.Number)
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
				} else if e.Status == elevator.Up {
					strings[k] += "↑"
				} else if e.Status == elevator.Down {
					strings[k] += "↓"
				} else {
					strings[k] += " "
				}
			} else {
				strings[k] += " "
			}
			strings[k] += "|"
		}
		elevInfo += " "
		elevInfo += strconv.Itoa(e.People)
	}

	for i, j := 0, len(keys)-1; i < j; i, j = i+1, j-1 {
		keys[i], keys[j] = keys[j], keys[i]
	}

	people := map[int]int{}
	for _, h := range f.Humans {
		if h.CurrentFloor == h.TargetFloor {
			if i, ok := people[h.CurrentFloor.Number]; ok {
				people[h.CurrentFloor.Number] = i + 1
			} else {
				people[h.CurrentFloor.Number] = 0
			}
		}
	}
	for _, k := range keys {
		strings[k] += " "
		strings[k] += strconv.Itoa(people[k])
		strings[k] += " "
	}
	for _, h := range f.Humans {
		if h.Elevator == nil && h.CurrentFloor != h.TargetFloor {
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
	out.WriteString("      ")
	out.WriteString(elevInfo)
	out.WriteString("\n")
	out.WriteString("\n")

	str := out.String()
	fmt.Print(str)
}
