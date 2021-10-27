package floor

type Floor struct {
	Number    int
	UpFloor   *Floor
	DownFloor *Floor
	Up        bool
	Down      bool
	Height    int
}

func New(number int, height int) *Floor {
	f := Floor{Number: number, Height: height, Up: false, Down: false}

	return &f
}

func NewArray(numbers []int, heights []int) []*Floor {
	var floors []*Floor

	for i, n := range numbers {
		f := New(n, heights[i])

		if i > 0 {
			f.DownFloor = floors[i-1]
			floors[i-1].UpFloor = f
		}
		floors = append(floors, f)
	}

	return floors
}

func (f *Floor) Compare(tf *Floor) int {
	if f.Number < tf.Number {
		return 1
	} else if f.Number == tf.Number {
		return 0
	} else {
		return -1
	}
}

func (f *Floor) IsHigh(height int) bool {
	return f.Height <= height
}
func (f *Floor) IsLow(height int) bool {
	return f.Height >= height
}

func (f *Floor) PushUp() {
	f.Up = true
}
func (f *Floor) PushDown() {
	f.Down = true
}
func (f *Floor) Arrive(status int) {
	switch status {
	case 1:
		f.Up = false
	case 0:
		f.Up = false
		f.Down = false
	case -1:
		f.Down = false
	}
}
