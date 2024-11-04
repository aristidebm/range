package main

import (
	"log"
	"math"
	"testing"
)

func TestFactoryMethodValue(t *testing.T) {
	tests := []struct {
		name     string
		expected float64
		actual   float64
	}{
		{
			name:     "Empty Range Lower Bound",
			expected: math.Inf(+1),
			actual:   Empty().Lower.Value,
		},
		{
			name:     "Empty Range Upper Bound",
			expected: math.Inf(-1),
			actual:   Empty().Upper.Value,
		},
		{
			name:     "Opened Range Lower Bound",
			expected: 10,
			actual:   Opened(10, 20).Lower.Value,
		},
		{
			name:     "Opened Range Upper Bound",
			expected: 20,
			actual:   Opened(10, 20).Upper.Value,
		},
		{
			name:     "Closed Range Lower Bound",
			expected: 10,
			actual:   Closed(10, 20).Lower.Value,
		},
		{
			name:     "Closed Range Upper Bound",
			expected: 20,
			actual:   Closed(10, 20).Upper.Value,
		},
		{
			name:     "ClosedOpened Range Lower Bound",
			expected: 10,
			actual:   ClosedOpened(10, 20).Lower.Value,
		},
		{
			name:     "ClosedOpened Range Upper Bound",
			expected: 20,
			actual:   ClosedOpened(10, 20).Upper.Value,
		},
		{
			name:     "OpenedClosed Range Lower Bound",
			expected: 10,
			actual:   OpenedClosed(10, 20).Lower.Value,
		},
		{
			name:     "OpenedClosed Range Upper Bound",
			expected: 20,
			actual:   OpenedClosed(10, 20).Upper.Value,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.actual != tt.expected {
				log.Fatalf("Expecting %v, but got %v", tt.expected, tt.actual)
			}
		})
	}
}
