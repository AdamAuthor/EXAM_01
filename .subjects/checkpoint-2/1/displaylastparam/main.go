package main

import (
	"academie/tester"
	"strings"

	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	table := append(random.StrSlice(chars.Words),
		" ",
		"1",
		"1 2",
		"1 2 3",
	)
	for _, s := range table {
		tester.Program(".subjects/checkpoint-2/1/displaylastparam/solutions", "displaylastparam", strings.Fields(s)...)
	}
}
