package main

import (
	bar "academie/model"
	"fmt"
	"os"

	_ "github.com/01-edu/go-tests/lib/challenge"
	_ "github.com/01-edu/z01"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	bar := bar.Progress()
	if _, err := tea.NewProgram(bar).Run(); err != nil {
		fmt.Println("Oh no!", err)
		os.Exit(1)
	}
}
