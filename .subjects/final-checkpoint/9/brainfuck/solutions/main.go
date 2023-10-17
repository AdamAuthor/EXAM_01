//  █████╗  ██████╗ █████╗ ██████╗ ███████╗███╗   ███╗██╗███████╗     ██████╗ ███╗   ██╗███████╗
// ██╔══██╗██╔════╝██╔══██╗██╔══██╗██╔════╝████╗ ████║██║██╔════╝    ██╔═══██╗████╗  ██║██╔════╝
// ███████║██║     ███████║██║  ██║█████╗  ██╔████╔██║██║█████╗      ██║   ██║██╔██╗ ██║█████╗
// ██╔══██║██║     ██╔══██║██║  ██║██╔══╝  ██║╚██╔╝██║██║██╔══╝      ██║   ██║██║╚██╗██║██╔══╝
// ██║  ██║╚██████╗██║  ██║██████╔╝███████╗██║ ╚═╝ ██║██║███████╗    ╚██████╔╝██║ ╚████║███████╗
// ╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝╚═════╝ ╚══════╝╚═╝     ╚═╝╚═╝╚══════╝     ╚═════╝ ╚═╝  ╚═══╝╚══════╝

package main

import (
	"fmt"
	"os"
)

func brainfuck(str string) {
	var ptr [2048]byte
	var tapePtr int

	for i := 0; i < len(str); i++ {
		switch str[i] {
		case '>':
			tapePtr++
		case '<':
			tapePtr--
		case '+':
			ptr[tapePtr]++
		case '-':
			ptr[tapePtr]--
		case '[':
			if ptr[tapePtr] == 0 {
				loopDepth := 1
				for loopDepth > 0 {
					i++
					if str[i] == '[' {
						loopDepth++
					} else if str[i] == ']' {
						loopDepth--
					}
				}
			}
		case ']':
			if ptr[tapePtr] != 0 {
				loopDepth := 1
				for loopDepth > 0 {
					i--
					if str[i] == ']' {
						loopDepth++
					} else if str[i] == '[' {
						loopDepth--
					}
				}
				i-- // Move back to the matching '['
			}
		case '.':
			fmt.Print(string(ptr[tapePtr]))
		}
	}
}

func main() {
	if len(os.Args) == 2 && len(os.Args[1]) > 0 {
		brainfuck(os.Args[1])
	}
	fmt.Println()
}
