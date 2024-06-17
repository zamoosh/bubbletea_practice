package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	columnWidth  = 50
	columnHeight = 20
)

type model struct {
	leftViewport  viewport.Model
	rightViewport viewport.Model
}

func initialModel() model {
	leftVP := viewport.New(columnWidth, columnHeight)
	rightVP := viewport.New(columnWidth, columnHeight)

	// Fill the viewports with some long text
	leftVP.SetContent(generateLongText("Left Column"))
	rightVP.SetContent(generateLongText("Right Column"))

	return model{
		leftViewport:  leftVP,
		rightViewport: rightVP,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := message.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "up":
			m.leftViewport.LineUp(1)
			m.rightViewport.LineUp(1)
		case "down":
			m.leftViewport.LineDown(1)
			m.rightViewport.LineDown(1)
		}
	case tea.MouseEvent:

		// Handle mouse scroll and drag events for both viewports
		switch msg.Type {
		// switch msg.Action {
		case tea.MouseWheelUp:
			m.leftViewport.LineUp(1)
			m.rightViewport.LineUp(1)
		case tea.MouseWheelDown:
			m.leftViewport.LineDown(1)
			m.rightViewport.LineDown(1)
		}
	}

	m.leftViewport, cmd = m.leftViewport.Update(message)
	m.rightViewport, cmd = m.rightViewport.Update(message)

	return m, cmd
}

func (m model) View() string {
	// Render the two viewports side by side
	leftColumn := m.leftViewport.View()
	rightColumn := m.rightViewport.View()

	// Combine the two columns
	return lipgloss.JoinHorizontal(lipgloss.Top, leftColumn, rightColumn)
}

func generateLongText(prefix string) string {
	var lines []string
	for i := 0; i < 100; i++ {
		lines = append(lines, fmt.Sprintf("[%02d] %s: This is line number %d \t", i, prefix, i+1))
	}
	return strings.Join(lines, "\n")
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
