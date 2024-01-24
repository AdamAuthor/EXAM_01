package main

import (
	"academie/tester"
	"strings"

	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	table := append(random.StrSlice(chars.Words),
		"dsfda",
		"",
		"1",
		"1",
	)
	for _, s := range table {
		tester.Program(".subjects/checkpoint-2/1/displayz/solutions", "displayz", strings.Fields(s)...)
	}
	tester.Program(".subjects/checkpoint-2/1/displayz/solutions", "displayz", "1", "z")
}
