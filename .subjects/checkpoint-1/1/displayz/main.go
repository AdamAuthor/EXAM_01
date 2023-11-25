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
		tester.Program(".subjects/checkpoint-1/1/displaya/solutions", "displaya", strings.Fields(s)...)
	}
	tester.Program(".subjects/checkpoint-1/1/displaya/solutions", "displaya", "1", "z")
}
