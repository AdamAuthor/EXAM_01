package main

import (
	"academie/tester"

	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	args := []string{
		"(johndoe)",
		")()",
		"([)]",
		"{2*[d - 3]/(12)}",
	}

	// 18 random tests ( at least half are valid)
	for i := 0; i < 3; i++ {
		args = append(args,
			"("+random.Str(chars.ASCII, 13)+")",
			"["+random.Str(chars.ASCII, 13)+"]",
			"{"+random.Str(chars.ASCII, 13)+"}",
			"("+random.Str(chars.Alnum, 13)+")",
			"["+random.Str(chars.Alnum, 13)+"]",
			"{"+random.Str(chars.Alnum, 13)+"}",
		)
	}

	tester.Program(".subjects/final-checkpoint/8/brackets/solutions", "brackets")

	for _, v := range args {
		tester.Program(".subjects/final-checkpoint/8/brackets/solutions", "brackets", v)
	}

	multArg := [][]string{
		{"", "{[(0 + 0)(1 + 1)](3*(-1)){()}}"},
		{"{][]}", "{3*[21/(12+ 23)]}"},
		{"{([)])}", "{{{something }- [something]}}", "there are"},
	}

	for _, v := range multArg {
		tester.Program(".subjects/final-checkpoint/8/brackets/solutions", "brackets", v...)
	}
}
