package ml

// Label represents data label type.
type Label struct {
	value int
}

var (
	Positive = &Label{value: 1}
	Negative = &Label{value: -1}
)

func (l *Label) IsPositive() bool {
	return l.value > 0
}

func (l *Label) IsNegative() bool {
	return l.value < 0
}
