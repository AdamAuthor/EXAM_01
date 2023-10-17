package exam

import (
	"fmt"
	"os"
	"strconv"
)

func AddTask(checkpoint string) map[int][]string {
	tasks := make(map[int][]string)
	// tasks[1] = []string{"countdown", "strlen"}
	// tasks[2] = []string{"firstrune", "lastrune", "lastword", "reduceint", "rot13"}
	// tasks[3] = []string{"alphamirror", "chunk", "tabmult"}
	// tasks[4] = []string{"inter", "ispowerof2", "piglatin", "romannumbers", "swapbits", "union"}
	// tasks[5] = []string{"printhex", "repeatalpha"}
	// tasks[6] = []string{"hiddenp", "revwstr", "rostring", "slice", "split"}
	// tasks[7] = []string{"addprimesum", "atoibase", "fprime", "itoa"}
	// tasks[8] = []string{"brackets", "gcd", "grouping", "listsize", "options", "printmemory", "rpncalc"}
	// tasks[9] = []string{"brainfuck", "itoabase", "listremoveif"}

	files, err := os.ReadDir("./.subjects/" + checkpoint)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	var exams []string

	for i := range files {
		taskNames, err := os.ReadDir("./.subjects/" + checkpoint + "/" + strconv.Itoa(i+1) + "/")
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}
		for _, name := range taskNames {
			exams = append(exams, string(name.Name()))
		}
		tasks[i+1] = exams
		exams = []string{}
	}

	return tasks
}
