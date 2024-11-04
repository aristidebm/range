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

func (r Range) Equal() bool {
	return true
}

func (r Range) IsEmpty() bool {
	return true
}

func (r Range) Difference(other Range) Range {
	return Range{}
}

func (r Range) Intersection(other Range) Range {
	return Range{}
}

func (r Range) Union(other Range) Range {
	return Range{}
}

func (r Range) Iter(step float64) float64 {
	return 0
}

func Empty() Range {
	return newRange(Bound{Value: math.Inf(+1), Type: LPAREN}, Bound{Value: math.Inf(-1), Type: RPAREN})
}

func OpenClosed(lower float64, upper float64) Range {
	return newRange(Bound{Value: lower, Type: LPAREN}, Bound{Value: upper, Type: RBRACKET})
}

func ClosedOpen(lower float64, upper float64) Range {
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
