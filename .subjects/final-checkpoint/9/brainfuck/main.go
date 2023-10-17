package main

import (
	"academie/tester"
	"strings"

	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	// individual tests 1)Hello World! 2)Hi 3)abc 4)ABC

	args := []string{
		"++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>.",
		"+++++[>++++[>++++H>+++++i<<-]>>>++\n<<<<-]>>--------.>+++++.>.",
		"++++++++++[>++++++++++>++++++++++>++++++++++<<<-]>---.>--.>-.>++++++++++.",
		"ld++++++++++++++++++++++++++++++++++++++++++++this+is++a++++comment++++++++++++++[>d+<-]>.+.+.>++++++++++.",

		strings.Join(
			[]string{
				"++++++++++[>++++++++++>++++++++++>++++++++++<<<-]>---",
				random.Str(".+", random.IntBetween(1, 25)),
				random.Str(".a", random.IntBetween(1, 100)),
				">--",
				random.Str(".+", random.IntBetween(1, 20)),
				random.Str(".a", random.IntBetween(1, 100)),
				">-",
				random.Str(".+", random.IntBetween(1, 20)),
				random.Str(".a", random.IntBetween(1, 100)),
				">++++++++++.",
			},
			"",
		),

		strings.Join(
			[]string{
				"ld++++++++++++++++++++++++++++++++++++++++++++this+is++a++++comment++++++++++++++[>d+<-]>.+",
				random.Str(".+", random.IntBetween(1, 10)),
				".+.>++++++++++.",
			},
			"",
		),
	}
	for _, v := range args {
		tester.Program(".subjects/final-checkpoint/9/brainfuck/solutions", "brainfuck", v)
	}
}
