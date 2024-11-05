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
	return r.lower.Value <= other.lower.Value &&
		other.lower.Value <= other.upper.Value &&
		other.upper.Value <= r.lower.Value
}

func (r interval) Belongs(other interval) bool {
	return other.lower.Value <= r.lower.Value &&
		r.lower.Value <= r.upper.Value &&
		r.upper.Value <= other.lower.Value
}

func (r interval) Equal(other interval) bool {
	return r.lower.Equal(other.lower) && r.upper.Equal(other.upper)
}

func (r interval) IsEmpty() bool {
	return r.lower.Value == math.Inf(+1) &&
		r.lower.Type == utils.LPAREN &&
		r.upper.Value == math.Inf(-1) &&
		r.upper.Type == utils.RPAREN
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
	case other.lower.Value < r.lower.Value && other.upper.Value < r.lower.Value:
		return Empty()
	case r.lower.Value < other.lower.Value && r.upper.Value < other.lower.Value:
		return Empty()
	default:
		lower := utils.Bound{Value: math.Max(r.lower.Value, other.lower.Value)}
		upper := utils.Bound{Value: math.Min(r.upper.Value, other.upper.Value)}

		if lower.Value == r.lower.Value {
			lower.Type = r.lower.Type
		} else {
			lower.Type = other.lower.Type
		}

		if upper.Value == r.upper.Value {
			upper.Type = r.upper.Type
		} else {
			upper.Type = other.upper.Type
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
	case other.lower.Value < r.lower.Value && other.upper.Value < r.lower.Value:
		return Empty()
	case r.lower.Value < other.lower.Value && r.upper.Value < other.lower.Value:
		return Empty()
	default:
		lower := utils.Bound{Value: math.Min(r.lower.Value, other.lower.Value)}
		upper := utils.Bound{Value: math.Max(r.upper.Value, other.upper.Value)}

		if lower.Value == r.lower.Value {
			lower.Type = r.lower.Type
		} else {
			lower.Type = other.lower.Type
		}

		if upper.Value == r.upper.Value {
			upper.Type = r.upper.Type
		} else {
			upper.Type = other.upper.Type
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
	return fmt.Sprintf("%v%v,%v%v", r.lower.Type, r.lower.Value, r.upper.Value, r.upper.Type)
}

type interval struct {
	lower utils.Bound
	upper utils.Bound
}

func (r interval) Lower() utils.Bound {
	return r.lower
}

func (r interval) Upper() utils.Bound {
	return r.upper
}

func newInterval(lower utils.Bound, upper utils.Bound) interval {
	if (lower.Value > upper.Value) ||
		(lower.Value == upper.Value && lower.Type != upper.Type) ||
		(lower.Value == upper.Value && lower.Type == utils.LPAREN) ||
		(lower.Type != utils.LPAREN && lower.Type != utils.LBRACKET) ||
		(upper.Type != utils.RPAREN && upper.Type != utils.RBRACKET) {
		return emptyInterval()
	}
	return interval{lower: lower, upper: upper}
}

func emptyInterval() interval {
	return interval{lower: utils.Bound{Value: math.Inf(+1), Type: utils.LPAREN}, upper: utils.Bound{Value: math.Inf(-1), Type: utils.RPAREN}}
}
