package main

import (
	"academie/tester"

	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	args := []string{
		"",
		"Hello",
		"World",
		"Home",
		"Theorem",
		"Choumi is the best cat",
		"abracadaba 01!",
		"abc",
		"MaTheMatiCs",
	}

	args = append(args, random.StrSlice(chars.Alnum)...)

	for _, v := range args {
		tester.Program(".subjects/final-checkpoint/5/repeatalpha/solutions", "repeatalpha", v)
	}
	tester.Program(".subjects/final-checkpoint/5/repeatalpha/solutions", "repeatalpha")
	tester.Program(".subjects/final-checkpoint/5/repeatalpha/solutions", "repeatalpha", "", "")
}
