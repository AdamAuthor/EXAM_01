package main

import (
	"academie/tester"

	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	args := []string{
		"abc   ",
		"Let there     be light",
		"     AkjhZ zLKIJz , 23y",
		"",
	}
	args = append(args, random.StrSlice(chars.Words)...)

	for _, arg := range args {
		tester.Program(".subjects/final-checkpoint/6/rostring/solutions", "rostring", arg)
	}
	tester.Program(".subjects/final-checkpoint/6/rostring/solutions", "rostring")
	tester.Program(".subjects/final-checkpoint/6/rostring/solutions", "rostring", "this", "is")
	tester.Program(".subjects/final-checkpoint/6/rostring/solutions", "rostring", "not", "good", "for  you")
}
