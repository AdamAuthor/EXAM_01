package main

import (
	"academie/model"
	"fmt"
	"os"

	_ "github.com/01-edu/go-tests/lib/challenge"
	_ "github.com/01-edu/z01"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	bar := model.Progress()

	if _, err := tea.NewProgram(bar).Run(); err != nil {
		fmt.Println("Oh no!", err)
		os.Exit(1)
	}
	model.Clear()
}
