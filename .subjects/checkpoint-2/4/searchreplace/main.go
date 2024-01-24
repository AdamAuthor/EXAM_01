package main

import (
	"academie/tester"

	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	type nodeTest struct {
		dataSearched    string
		letterLookedFor string
		letterReplacing string
	}

	args := []nodeTest{
		{
			dataSearched:    "hélla",
			letterLookedFor: "é",
			letterReplacing: "o",
		},
		{
			dataSearched:    "hella",
			letterLookedFor: "z",
			letterReplacing: "o",
		},
		{
			dataSearched:    "hella",
			letterLookedFor: "h",
			letterReplacing: "o",
		},
	}

	for i := 0; i < 20; i++ {
		args = append(args, nodeTest{
			dataSearched:    random.Str(chars.Words, 13),
			letterLookedFor: random.Str(chars.Alnum, 1),
			letterReplacing: random.Str(chars.Alnum, 1),
		})
	}

	for _, arg := range args {
		tester.Program(".subjects/checkpoint-2/4/searchreplace/solutions", "searchreplace", arg.dataSearched, arg.letterLookedFor, arg.letterReplacing)
	}
}
