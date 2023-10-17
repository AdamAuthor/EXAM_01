package main

import (
	"academie/tester"

	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	table := append(random.StrSlice(chars.Words),
		" ",
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPRSTUVWXYZ",
		"a b c d e f g h ijklmnopqrstuvwxyz A B C D E FGHIJKLMNOPRSTUVWXYZ",
	)
	for _, s := range table {
		tester.Program(".subjects/final-checkpoint/2/rot13/solutions", "rot13", s)
	}
	tester.Program(".subjects/final-checkpoint/2/rot13/solutions", "rot13", "1 argument", "2 arguments")
	tester.Program(".subjects/final-checkpoint/2/rot13/solutions", "rot13", "1 argument", "2 arguments", "3 arguments")
}
