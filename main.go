package main

import (
	"fmt"
	"math"
)

type BoundType string

const (
	LPAREN   BoundType = "("
	RPAREN             = ")"
	LBRACKET           = "["
	RBRACKET           = "]"
)

type Bound struct {
	Value float64
	Type  BoundType
}

func (b Bound) Equal(other Bound) bool {
	return b.Value == other.Value && b.Type == other.Type
}

type Range struct {
	Lower Bound
	Upper Bound
}

func (r Range) String() string {
	if r.IsEmpty() {
		return fmt.Sprintf("%s%s", LPAREN, RPAREN)
	}
	return fmt.Sprintf("%v%v,%v%v", r.Lower.Type, r.Lower.Value, r.Upper.Value, r.Upper.Type)
}

func (r Range) Contains(other Range) bool {
	return true
}

func (r Range) Belongs(other Range) bool {
	return true
}

func (r Range) Equal(other Range) bool {
	return r.Lower.Equal(other.Lower) && r.Upper.Equal(other.Upper)
}

func (r Range) IsEmpty() bool {
	return r.Lower.Value == math.Inf(+1) &&
		r.Lower.Type == LPAREN &&
		r.Upper.Value == math.Inf(-1) &&
		r.Upper.Type == RPAREN
}

func (r Range) Difference(other Range) Range {
	if other.IsEmpty() {
		// structs are passed by value, no need to clone
		return other
	}
	return Range{}
}

func (r Range) Intersection(other Range) Range {
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

func (r Range) intersection(other Range) Range {
	switch {
	// They are disjoints
	case other.Lower.Value < r.Lower.Value && other.Upper.Value < r.Lower.Value:
		return Empty()
	case r.Lower.Value < other.Lower.Value && r.Upper.Value < other.Lower.Value:
		return Empty()
	default:
		lower := Bound{Value: math.Max(r.Lower.Value, other.Lower.Value)}
		upper := Bound{Value: math.Min(r.Upper.Value, other.Upper.Value)}

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
		return newRange(lower, upper)
	}
}

func (r Range) Union(other Range) Range {
	if other.IsEmpty() {
		// structs are passed by value, no need to clone
		return r
	}
	return Range{}
}

func (r Range) Iter(step float64) float64 {
	return 0
}

func Empty() Range {
	return newRange(Bound{Value: math.Inf(+1), Type: LPAREN}, Bound{Value: math.Inf(-1), Type: RPAREN})
}

func OpenedClosed(lower float64, upper float64) Range {
	return newRange(Bound{Value: lower, Type: LPAREN}, Bound{Value: upper, Type: RBRACKET})
}

func ClosedOpened(lower float64, upper float64) Range {
	return newRange(Bound{Value: lower, Type: LBRACKET}, Bound{Value: upper, Type: RPAREN})
}

func Opened(lower float64, upper float64) Range {
	return newRange(Bound{Value: lower, Type: LPAREN}, Bound{Value: upper, Type: RPAREN})
}

func Closed(lower float64, upper float64) Range {
	return newRange(Bound{Value: lower, Type: LBRACKET}, Bound{Value: upper, Type: RBRACKET})
}

func newRange(lower Bound, upper Bound) Range {
	return Range{
		Lower: lower, Upper: upper,
	}
}
