package main

import (
	"academie/tester"

	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	args := []string{
		"", "pig", "is", "crunch", "crnch", "something else",
	}

	for i := 0; i < 4; i++ {
		args = append(args, random.Str(chars.Basic, 7))
	}

	for _, v := range args {
		tester.Program(".subjects/final-checkpoint/4/piglatin/solutions", "piglatin", v)
	}
}
