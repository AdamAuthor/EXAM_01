//  █████╗  ██████╗ █████╗ ██████╗ ███████╗███╗   ███╗██╗███████╗     ██████╗ ███╗   ██╗███████╗
// ██╔══██╗██╔════╝██╔══██╗██╔══██╗██╔════╝████╗ ████║██║██╔════╝    ██╔═══██╗████╗  ██║██╔════╝
// ███████║██║     ███████║██║  ██║█████╗  ██╔████╔██║██║█████╗      ██║   ██║██╔██╗ ██║█████╗
// ██╔══██║██║     ██╔══██║██║  ██║██╔══╝  ██║╚██╔╝██║██║██╔══╝      ██║   ██║██║╚██╗██║██╔══╝
// ██║  ██║╚██████╗██║  ██║██████╔╝███████╗██║ ╚═╝ ██║██║███████╗    ╚██████╔╝██║ ╚████║███████╗
// ╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝╚═════╝ ╚══════╝╚═╝     ╚═╝╚═╝╚══════╝     ╚═════╝ ╚═╝  ╚═══╝╚══════╝

package exam

import (
	dp "academie/exam/displayerror"
	"academie/model"
	"academie/model/congrats"
	"fmt"
	"os"
	"os/exec"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

var access = false
var exit = false

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render

type example struct {
	viewport viewport.Model
}

func createGlamour(content string) (*example, error) {
	const width = 78

	vp := viewport.New(width, 20)
	vp.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		PaddingRight(2)

	renderer, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(width),
	)
	if err != nil {
		return nil, err
	}

	str, err := renderer.Render(content)
	if err != nil {
		return nil, err
	}

	vp.SetContent(str)

	return &example{
		viewport: vp,
	}, nil
}

func (e example) Init() tea.Cmd {
	return nil
}

func (e example) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			exit = true
			exec.Command("reset")
			return e, tea.Quit
		case "enter":
			access = true
			return e, tea.Quit
		default:
			var cmd tea.Cmd
			e.viewport, cmd = e.viewport.Update(msg)
			return e, cmd
		}
	default:
		return e, nil
	}
}

func (e example) View() string {
	return e.viewport.View() + e.helpView()
}

func (e example) helpView() string {
	return helpStyle("\n  ↑/↓: Navigate • Enter: Start Testing • q: Exit\n")
}

func (c *Checkpoint) InitTask() int {
	if c.Level > c.LenLevels {
		model.Clear()
		congrats.Congrats()
		os.Exit(0)
	}
	if exit {
		model.Clear()
		os.Exit(0)
	}
	var content, typeCheckpoint string
	content, c.Task, typeCheckpoint = OpenFile(c.IdxCheck, c.Level, c.Task)
	bModel, err := createGlamour(content)
	if err != nil {
		fmt.Println("Could not initialize Bubble Tea model:", err)
		os.Exit(1)
	}

	if _, err := tea.NewProgram(bModel).Run(); err != nil {
		fmt.Println("Bummer, there's been an error:", err)
		os.Exit(1)
	}
	var strErr string
	if !exit {
		c.Level, strErr = RunTester(c.Task, c.Level, typeCheckpoint)
	}
	if strErr != "" {
		model.Clear()
		dp.DisplayError(strErr)
		model.Clear()
		return c.InitTask()
	}
	model.Clear()
	c.Task = ""
	return c.InitTask()
}
