package floor

type Floor struct {
	number    int
	upFloor   *Floor
	downFloor *Floor
	up        bool
	down      bool
	height    int
}

func New(number int, height int) *Floor {
	f := Floor{number: number, height: height, up: false, down: false}

	return &f
}

func NewArray(numbers []int, heights []int) []*Floor {
	var floors []*Floor

	for i, n := range numbers {
		f := New(n, heights[i])

		if i > 0 {
			f.downFloor = floors[i-1]
			floors[i-1].upFloor = f
		}
		floors = append(floors, f)
	}

	return floors
}

func (f *Floor) Compare(tf *Floor) bool {
	return f.number < tf.number
}

func (f *Floor) PushUp()    { f.up = true }
func (f *Floor) PushDown()  { f.down = true }
func (f *Floor) Up() bool   { return f.up }
func (f *Floor) Down() bool { return f.down }
