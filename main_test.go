package main

import (
	"fmt"
	"log"
	"math"
	"testing"
)

func TestFactoryMethod(t *testing.T) {
	emp := Empty()
	if emp.Lower.Value != math.Inf(+1) {
		log.Fatal(fmt.Sprintf("name=%v actual=%v expecting=%v", "Empty set lower bound", emp.Lower.Value, math.Inf(+1)))
	}
}
