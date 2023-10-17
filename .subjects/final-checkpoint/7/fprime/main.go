package main

import (
	"academie/tester"
	"strconv"

	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	table := []string{
		" ",
		"1",
		"1 1",
		"hello",
		"p 1",
		"804577",
		"225225",
		"8333325",
		"42",
		"9539",
		"1000002",
		"1000003",
	}
	for i := 0; i < 10; i++ {
		table = append(table, strconv.Itoa(random.IntBetween(1, 100)))
	}
	for _, s := range table {
		tester.Program(".subjects/final-checkpoint/7/fprime/solutions", "fprime", s)
	}
}
