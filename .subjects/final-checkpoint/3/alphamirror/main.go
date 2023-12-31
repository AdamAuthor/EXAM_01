package main

import (
	"academie/tester"

	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	args := [][]string{
		{""},
		{"One", "ring!"},
		{"testing spaces and #!*"},
		{"more", "than", "three", "arguments"},
		{"Upper anD LoWer cAsE"},
		{random.Str(chars.Words, 13)},
	}

	for _, v := range args {
		tester.Program(".subjects/final-checkpoint/3/alphamirror/solutions", "alphamirror", v...)
	}
}
