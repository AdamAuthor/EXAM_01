package main

import (
	"academie/tester"
	"strconv"
	"strings"

	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	table := []string{
		" ",
		"123 132 1",
		"1 5",
		"0",
	}
	for i := 0; i < 10; i++ {
		table = append(table, strconv.Itoa(random.IntBetween(-1000, random.MaxInt)))
	}
	for i := 0; i < 15; i++ {
		table = append(table, strconv.Itoa(i))
	}
	for _, s := range table {
		tester.Program(".subjects/final-checkpoint/5/printhex/solutions", "printhex", strings.Fields(s)...)
	}
}
