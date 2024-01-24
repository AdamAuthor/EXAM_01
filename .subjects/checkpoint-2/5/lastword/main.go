package main

import (
	"academie/tester"

	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	args := []string{
		"FOR PONY",
		"this ... is sparta, then again, maybe not",
		" ",
		" lorem,ipsum ",
	}

	args = append(args, random.StrSlice(chars.Words)...)

	for _, v := range args {
		tester.Program(".subjects/checkpoint-2/5/lastword/solutions", "lastword", v)
	}

	tester.Program(".subjects/checkpoint-2/5/lastword/solutions", "lastword", "a", "b")
	tester.Program(".subjects/checkpoint-2/5/lastword/solutions", "lastword")
}
