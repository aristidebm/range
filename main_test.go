package main

import (
	"log"
	"math"
	"testing"
)

func TestEqual(t *testing.T) {
	emp1 := Empty()
	emp2 := Empty()
	if !emp1.Equal(emp2) {
		log.Fatalf("actual: %v != %v; expected: %v == %v", emp1, emp2, emp1, emp2)
	}
}

func TestFactoryMethod(t *testing.T) {
	tests := []struct {
		name     string
		expected Range
		actual   Range
	}{
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.actual.Equal(tt.expected) {
				log.Fatalf("actual: %v; expected: %v", tt.expected, tt.actual)
			}
		})
	}
}
