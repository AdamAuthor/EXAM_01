//  █████╗  ██████╗ █████╗ ██████╗ ███████╗███╗   ███╗██╗███████╗     ██████╗ ███╗   ██╗███████╗
// ██╔══██╗██╔════╝██╔══██╗██╔══██╗██╔════╝████╗ ████║██║██╔════╝    ██╔═══██╗████╗  ██║██╔════╝
// ███████║██║     ███████║██║  ██║█████╗  ██╔████╔██║██║█████╗      ██║   ██║██╔██╗ ██║█████╗
// ██╔══██║██║     ██╔══██║██║  ██║██╔══╝  ██║╚██╔╝██║██║██╔══╝      ██║   ██║██║╚██╗██║██╔══╝
// ██║  ██║╚██████╗██║  ██║██████╔╝███████╗██║ ╚═╝ ██║██║███████╗    ╚██████╔╝██║ ╚████║███████╗
// ╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝╚═════╝ ╚══════╝╚═╝     ╚═╝╚═╝╚══════╝     ╚═════╝ ╚═╝  ╚═══╝╚══════╝

package main

import (
	"academie/tester"
	"strconv"

	"github.com/01-edu/go-tests/lib/is"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	// adds random numbers
	args := random.IntSliceBetween(1, 10000)

	// fill with all prime numbers between 0 and 100
	for i := 0; i < 100; i++ {
		if is.Prime(i) {
			args = append(args, i)
		}
	}
	for _, i := range args {
		tester.Program(".subjects/final-checkpoint/7/addprimesum/solutions", "addprimesum", strconv.Itoa(i))
	}
	// special cases
	tester.Program(".subjects/final-checkpoint/7/addprimesum/solutions", "addprimesum")
	tester.Program(".subjects/final-checkpoint/7/addprimesum/solutions", "addprimesum", `""`)
	tester.Program(".subjects/final-checkpoint/7/addprimesum/solutions", "addprimesum", "1", "2")
}
