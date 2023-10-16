package main

import (
	"academie/model"
	"fmt"

	_ "github.com/01-edu/go-tests/lib/challenge"
	_ "github.com/01-edu/z01"
)

func main() {
	model.RunProgress()
	access := model.RunTerms()

	if access {
		fmt.Println("Hello World!")
	} else {
		return
	}
}
