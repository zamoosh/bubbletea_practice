package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	boxWidth  = 50
	boxHeight = 10
)

type model struct {
	viewport viewport.Model
	content  string
}

func initialModel() model {
	vp := viewport.New(boxWidth, boxHeight)
	vp.MouseWheelEnabled = true // Enable mouse wheel scrolling

	// Generate some long text for the viewport content
	content := generateLongText()
	vp.SetContent(content)

	return model{
		viewport: vp,
		content:  content,
	}
}

func (m model) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (m model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := message.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	case tea.MouseMsg:
		switch msg.Button {
		case tea.MouseButtonWheelUp:
			fmt.Println("mouse wheel up!")
		default:
			panic("unhandled default case")

		}
	}

	m.viewport, cmd = m.viewport.Update(message)
	return m, cmd
}

func (m model) View() string {
	scrollbar := renderScrollbar(m.viewport)
	content := m.viewport.View()

	return lipgloss.NewStyle().
		Width(boxWidth + 3). // Extra space for the scrollbar
		Height(boxHeight).
		Border(lipgloss.RoundedBorder()).
		Render(lipgloss.JoinHorizontal(lipgloss.Left, content, scrollbar))
}

func renderScrollbar(vp viewport.Model) string {
	visibleLines := vp.Height
	// totalLines := len(strings.Split(vp.View(), "\n"))
	viewportLines := vp.YOffset + visibleLines
	scrollbar := ""

	for i := 0; i < visibleLines; i++ {
		if i < vp.YOffset {
			scrollbar += " "
		} else if i < viewportLines {
			scrollbar += "█"
		} else {
			scrollbar += " "
		}
	}

	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FF00FF")).
		Render(scrollbar)
}

func generateLongText() string {
	lines := []string{}
	for i := 0; i < 100; i++ {
		lines = append(lines, fmt.Sprintf("This is line number %d", i+1))
	}
	return strings.Join(lines, "\n")
}

func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
