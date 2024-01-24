package main

import (
	"academie/tester"

	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	table := append(random.StrSlice(chars.Words),
		"",
		"             a as",
		"   f     d",
		"   asd    ad",
		"   salut !!! ",
		" salut ! ! !",
		"salut ! !",
	)
	for _, s := range table {
		tester.Program(".subjects/checkpoint-2/5/firstword/solutions", "firstword", s)
	}
}
