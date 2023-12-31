package main

import (
	"academie/tester"

	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	args := []string{
		"",
		"abcdefghijklm",
		"the time of contempt precedes that of indifference",
		"he stared at the mountain",
		"qw qw e qwsa d",
	}

	// 3 valid random sentences with no spaces at the beginning nor the end and only one space for separator.
	for i := 0; i < 3; i++ {
		numberOfWords := random.IntBetween(1, 6)
		sentence := random.Str(chars.Alnum, 13)
		for j := 0; j < numberOfWords; j++ {
			sentence += " " + random.Str(chars.Alnum, 13)
		}
		sentence += random.Str(chars.Alnum, 13)
		args = append(args, sentence)
	}

	for _, s := range args {
		tester.Program(".subjects/final-checkpoint/6/revwstr/solutions", "revwstr", s)
	}

	tester.Program(".subjects/final-checkpoint/6/revwstr/solutions", "revwstr")
	tester.Program(".subjects/final-checkpoint/6/revwstr/solutions", "revwstr", "1param", "2param", "3param", "4param")
}
