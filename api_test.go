package ranges

import (
	"log"
	"math"
	"testing"

	"example.com/ranges/internal/utils"
)

type TableEntry struct {
	name     string
	expected interval
	actual   interval
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
			expected: interval{lower: utils.Bound{Value: math.Inf(+1), Type: utils.LPAREN}, upper: utils.Bound{Value: math.Inf(-1), Type: utils.RPAREN}},
			actual:   Empty(),
		},
		{
			name:     "Opened",
			expected: interval{lower: utils.Bound{Value: 10, Type: utils.LPAREN}, upper: utils.Bound{Value: 20, Type: utils.RPAREN}},
			actual:   Opened(10, 20),
		},
		{
			name:     "Closed",
			expected: interval{lower: utils.Bound{Value: 10, Type: utils.LBRACKET}, upper: utils.Bound{Value: 20, Type: utils.RBRACKET}},
			actual:   Closed(10, 20),
		},
		{
			name:     "ClosedOpened",
			expected: interval{lower: utils.Bound{Value: 10, Type: utils.LBRACKET}, upper: utils.Bound{Value: 20, Type: utils.RPAREN}},
			actual:   ClosedOpened(10, 20),
		},
		{
			name:     "OpenedClosed",
			expected: interval{lower: utils.Bound{Value: 10, Type: utils.LPAREN}, upper: utils.Bound{Value: 20, Type: utils.RBRACKET}},
			actual:   OpenedClosed(10, 20),
		},
	}
	assertTable(t, table)
}

func TestInvalidBounds(t *testing.T) {
	table := []TableEntry{
		{
			name:     "lowerutils.BoundGreatherThanupperutils.Bound",
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
