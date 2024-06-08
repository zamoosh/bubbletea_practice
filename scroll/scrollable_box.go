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
	boxHeight = 20
)

type model struct {
	viewport viewport.Model
}

func initialModel() model {
	longText := generateLongText()

	vp := viewport.New(boxWidth, boxHeight)
	vp.SetContent(longText)

	return model{
		viewport: vp,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

func (m model) View() string {
	boxStyle := lipgloss.NewStyle().
		Width(boxWidth).
		Height(boxHeight).
		Border(lipgloss.RoundedBorder()).
		Padding(1, 2)

	return boxStyle.Render(m.viewport.View())
}

func generateLongText() string {
	var lines []string
	for i := 0; i < 400; i++ {
		lines = append(lines, fmt.Sprintf("This is line number %d", i+1))
	}
	return strings.Join(lines, "\n")
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
