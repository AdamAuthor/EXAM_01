package main

import (
	"academie/tester"
	"strconv"

	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	var table []string
	for i := 1; i < 10; i++ {
		table = append(table, strconv.Itoa(i))
	}
	for i := 0; i < 5; i++ {
		table = append(table, strconv.Itoa(random.IntBetween(1, 1000)))
	}

	for _, arg := range table {
		tester.Program(".subjects/final-checkpoint/3/tabmult/solutions", "tabmult", arg)
	}
}
