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
        index := model.RunSelect()
        fmt.Println(index + 1)
    } else {
        return
    }
}
