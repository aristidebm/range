package main

import (
	"log"
	"math"
	"testing"
)

type TableEntry struct {
	name     string
	expected Range
	actual   Range
}

func TestEqual(t *testing.T) {
	emp1 := Empty()
	emp2 := Empty()
	if !emp1.Equal(emp2) {
		log.Fatalf("actual: %v != %v; expected: %v == %v", emp1, emp2, emp1, emp2)
	}
}

func TestFactoryMethod(t *testing.T) {
	table := []TableEntry{
		{
			name:     "Empty",
			expected: Range{Lower: Bound{Value: math.Inf(+1), Type: LPAREN}, Upper: Bound{Value: math.Inf(-1), Type: RPAREN}},
			actual:   Empty(),
		},
		{
			name:     "Opened",
			expected: Range{Lower: Bound{Value: 10, Type: LPAREN}, Upper: Bound{Value: 20, Type: RPAREN}},
			actual:   Opened(10, 20),
		},
		{
			name:     "Closed",
			expected: Range{Lower: Bound{Value: 10, Type: LBRACKET}, Upper: Bound{Value: 20, Type: RBRACKET}},
			actual:   Closed(10, 20),
		},
		{
			name:     "ClosedOpened",
			expected: Range{Lower: Bound{Value: 10, Type: LBRACKET}, Upper: Bound{Value: 20, Type: RPAREN}},
			actual:   ClosedOpened(10, 20),
		},
		{
			name:     "OpenedClosed",
			expected: Range{Lower: Bound{Value: 10, Type: LPAREN}, Upper: Bound{Value: 20, Type: RBRACKET}},
			actual:   OpenedClosed(10, 20),
		},
	}
	assertTable(t, table)
}

func TestInvalidBounds(t *testing.T) {
	table := []TableEntry{
		{
			name:     "LowerBoundGreatherThanUpperBound",
			expected: Empty(),
			actual:   OpenedClosed(20, 10),
		},
	}
	assertTable(t, table)
}

func TestIntersection(t *testing.T) {
	table := []TableEntry{
		{
			name:     "EmptyAndEmpty",
			expected: Empty(),
			actual:   Empty().Intersection(Empty()),
		},
		{
			name:     "EmptyAndOpened",
			expected: Empty(),
			actual:   Empty().Intersection(Opened(10, 20)),
		},
		{
			name:     "EmptyAndClosed",
			expected: Empty(),
			actual:   Empty().Intersection(Closed(10, 20)),
		},
		{
			name:     "EmptyAndOpenedClosed",
			expected: Empty(),
			actual:   Empty().Intersection(OpenedClosed(10, 20)),
		},
		{
			name:     "EmptyAndClosedOpened",
			expected: Empty(),
			actual:   Empty().Intersection(ClosedOpened(10, 20)),
		},
		{
			name:     "Contains",
			expected: Opened(10, 20),
			actual:   ClosedOpened(5, 20).Intersection(Opened(10, 20)),
		},
		{
			name:     "Belongs",
			expected: ClosedOpened(10, 20),
			actual:   ClosedOpened(10, 20).Intersection(Opened(5, 20)),
		},
	}
	assertTable(t, table)
}

func TestUnion(t *testing.T) {
	table := []TableEntry{
		{
			name:     "EmptyOrEmpty",
			expected: Empty(),
			actual:   Empty().Union(Empty()),
		},
		{
			name:     "EmptyOrOpened",
			expected: Opened(10, 20),
			actual:   Empty().Union(Opened(10, 20)),
		},
		{
			name:     "EmptyOrClosed",
			expected: Closed(10, 20),
			actual:   Empty().Union(Closed(10, 20)),
		},
		{
			name:     "EmptyOrOpenedClosed",
			expected: OpenedClosed(10, 20),
			actual:   Empty().Union(OpenedClosed(10, 20)),
		},
		{
			name:     "EmptyOrClosedOpened",
			expected: ClosedOpened(10, 20),
			actual:   Empty().Union(ClosedOpened(10, 20)),
		},
	}
	assertTable(t, table)
}

func assertTable(t *testing.T, table []TableEntry) {
	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.actual.Equal(tt.expected) {
				log.Fatalf("actual: %v; expected: %v", tt.actual, tt.expected)
			}
		})
	}
}
