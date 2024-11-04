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
		log.Fatal(fmt.Sprintf("name=%v actual=%v expecting=%v", "Empty set lower bound value", emp.Lower.Value, math.Inf(+1)))
	}

	if emp.Upper.Value != math.Inf(-1) {
		log.Fatal(fmt.Sprintf("name=%v actual=%v expecting=%v", "Empty set upper bound value", emp.Upper.Value, math.Inf(-1)))
	}

	if emp.Lower.Type != LPAREN {
		log.Fatal(fmt.Sprintf("name=%v actual=%v expecting=%v", "Empty set lower bound type", emp.Lower.Type, LPAREN))
	}

	if emp.Upper.Type != RPAREN {
		log.Fatal(fmt.Sprintf("name=%v actual=%v expecting=%v", "Empty set upper bound type", emp.Upper.Type, RPAREN))
	}
}
