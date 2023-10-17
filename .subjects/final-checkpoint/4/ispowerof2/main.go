package main

import (
	"academie/tester"
	"strconv"

	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	args := [][]string{
		{"1"},
		{"2"},
		{"3"},
		{"4"},
		{"1024"},
		{"4096"},
		{"8388608"},
		{"1", "2"},
		{},
	}
	for i := 0; i < 12; i++ {
		args = append(args, []string{strconv.Itoa(random.IntBetween(1, 2048))})
	}
	for _, v := range args {
		tester.Program(".subjects/final-checkpoint/4/ispowerof2/solutions", "ispowerof2", v...)
	}
}
