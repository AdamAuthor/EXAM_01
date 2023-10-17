package main

import (
	"academie/tester"
	"strconv"

	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	table := []string{
		"0",
		"4000",
		"5000",
		"12433",
		"hello",
		"good luck",
	}
	for i := 0; i < 7; i++ {
		table = append(table, strconv.Itoa(random.IntBetween(0, 4000)))
	}
	for _, v := range table {
		tester.Program(".subjects/final-checkpoint/4/romannumbers/solutions", "romannumbers", v)
	}
}
