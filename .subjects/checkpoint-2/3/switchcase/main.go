package main

import (
	"academie/tester"

	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	table := append(random.StrSlice(chars.Words),
		" ",
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"abcdefghi jklmnop qrstuvwxyz ABCDEFGHI JKLMNOPQR STUVWXYZ ! ",
	)
	for _, s := range table {
		tester.Program(".subjects/checkpoint-2/3/switchcase/solutions", "switchcase", s)
	}
}
