package ranges

import (
	"example.com/ranges/internal/utils"
	"fmt"
	"math"
)

func (r interval) Contains(other interval) bool {
	if other.IsEmpty() {
		return true
	}
	return r.Lower.Value <= other.Lower.Value &&
		other.Lower.Value <= other.Upper.Value &&
		other.Upper.Value <= r.Lower.Value
}

func (r interval) Belongs(other interval) bool {
	return other.Lower.Value <= r.Lower.Value &&
		r.Lower.Value <= r.Upper.Value &&
		r.Upper.Value <= other.Lower.Value
}

func (r interval) Equal(other interval) bool {
	return r.Lower.Equal(other.Lower) && r.Upper.Equal(other.Upper)
}

func (r interval) IsEmpty() bool {
	return r.Lower.Value == math.Inf(+1) &&
		r.Lower.Type == utils.LPAREN &&
		r.Upper.Value == math.Inf(-1) &&
		r.Upper.Type == utils.RPAREN
}

func (r interval) Difference(other interval) interval {
	if other.IsEmpty() {
		return other
	}
	return interval{}
}

func (r interval) Intersection(other interval) interval {
	switch {
	case other.IsEmpty() || r.Contains(other):
		return other
	case r.IsEmpty() || other.Contains(r):
		return r
	case r.Equal(other):
		return r
	default:
		return r.intersection(other)
	}
}

func (r interval) intersection(other interval) interval {
	switch {
	// disjoin
	case other.Lower.Value < r.Lower.Value && other.Upper.Value < r.Lower.Value:
		return Empty()
	case r.Lower.Value < other.Lower.Value && r.Upper.Value < other.Lower.Value:
		return Empty()
	default:
		lower := utils.Bound{Value: math.Max(r.Lower.Value, other.Lower.Value)}
		upper := utils.Bound{Value: math.Min(r.Upper.Value, other.Upper.Value)}

		if lower.Value == r.Lower.Value {
			lower.Type = r.Lower.Type
		} else {
			lower.Type = other.Lower.Type
		}

		if upper.Value == r.Upper.Value {
			upper.Type = r.Upper.Type
		} else {
			upper.Type = other.Upper.Type
		}
		return newInterval(lower, upper)
	}
}

func (r interval) Union(other interval) interval {
	switch {
	case other.IsEmpty() || r.Contains(other):
		return r
	case r.IsEmpty() || other.Contains(r):
		return other
	case r.Equal(other):
		return r
	default:
		return r.union(other)
	}
}

func (r interval) union(other interval) interval {
	switch {
	// disjoin (by convention, we will consider union of disjoin range to be empty)
	case other.Lower.Value < r.Lower.Value && other.Upper.Value < r.Lower.Value:
		return Empty()
	case r.Lower.Value < other.Lower.Value && r.Upper.Value < other.Lower.Value:
		return Empty()
	default:
		lower := utils.Bound{Value: math.Min(r.Lower.Value, other.Lower.Value)}
		upper := utils.Bound{Value: math.Max(r.Upper.Value, other.Upper.Value)}

		if lower.Value == r.Lower.Value {
			lower.Type = r.Lower.Type
		} else {
			lower.Type = other.Lower.Type
		}

		if upper.Value == r.Upper.Value {
			upper.Type = r.Upper.Type
		} else {
			upper.Type = other.Upper.Type
		}
		return newInterval(lower, upper)
	}
}

func (r interval) Iter(step float64) float64 {
	return 0
}

func Empty() interval {
	return emptyInterval()
}

func OpenedClosed(lower float64, upper float64) interval {
	return newInterval(utils.Bound{Value: lower, Type: utils.LPAREN}, utils.Bound{Value: upper, Type: utils.RBRACKET})
}

func ClosedOpened(lower float64, upper float64) interval {
	return newInterval(utils.Bound{Value: lower, Type: utils.LBRACKET}, utils.Bound{Value: upper, Type: utils.RPAREN})
}

func Opened(lower float64, upper float64) interval {
	return newInterval(utils.Bound{Value: lower, Type: utils.LPAREN}, utils.Bound{Value: upper, Type: utils.RPAREN})
}

func Closed(lower float64, upper float64) interval {
	return newInterval(utils.Bound{Value: lower, Type: utils.LBRACKET}, utils.Bound{Value: upper, Type: utils.RBRACKET})
}

func (r interval) String() string {
	if r.IsEmpty() {
		return fmt.Sprintf("%s%s", utils.LPAREN, utils.RPAREN)
	}
	return fmt.Sprintf("%v%v,%v%v", r.Lower.Type, r.Lower.Value, r.Upper.Value, r.Upper.Type)
}

type interval struct {
	Lower utils.Bound
	Upper utils.Bound
}

func newInterval(lower utils.Bound, upper utils.Bound) interval {
	if (lower.Value > upper.Value) ||
		(lower.Value == upper.Value && lower.Type != upper.Type) ||
		(lower.Value == upper.Value && lower.Type == utils.LPAREN) ||
		(lower.Type != utils.LPAREN && lower.Type != utils.LBRACKET) ||
		(upper.Type != utils.RPAREN && upper.Type != utils.RBRACKET) {
		return emptyInterval()
	}
	return interval{Lower: lower, Upper: upper}
}

func emptyInterval() interval {
	return interval{Lower: utils.Bound{Value: math.Inf(+1), Type: utils.LPAREN}, Upper: utils.Bound{Value: math.Inf(-1), Type: utils.RPAREN}}
}
