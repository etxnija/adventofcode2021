package day1

//go:generate stringer -type=OneToTen
type OneToTen int

const (
	one OneToTen = iota + 1
	two
	three
	four
	five
	six
	seven
	eight
	nine
)

var numbers = []OneToTen{one, two, three, four, five, six, seven, eight, nine}

func (n OneToTen) String() string {
	switch n {
	case one:
		return "one"
	case two:
		return "two"
	case three:
		return "three"
	case four:
		return "four"
	case five:
		return "five"
	case six:
		return "six"
	case seven:
		return "seven"
	case eight:
		return "eight"
	case nine:
		return "nine"
	default:
		return ""
	}
}
