package main

import (
	"academie/exam"
	"academie/model"
	"fmt"

	_ "github.com/01-edu/go-tests/lib/challenge"
	_ "github.com/01-edu/z01"
)

func main() {
	model.Clear()
	model.RunProgress()
	access := model.RunTerms()
	var index int
	if access {
		index = model.RunSelect() + 1
	} else {
		return
	}
	testStart := exam.InitTask(index)

	if testStart {
		fmt.Println("Start testing...")
	}
}
