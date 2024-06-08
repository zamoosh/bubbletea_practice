package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	message string
}

func initialModel() model {
	return model{
		message: "Press Ctrl+Shift+B to trigger the event.",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		m.message = msg.String()

		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "alt+B":
			m.message = "ali ali ali!"
		}

		// if msg.String() == "B" && msg.Ctrl && msg.Shift {
		// 	m.message = "Ctrl+Shift+B was pressed!"
		// } else if msg.String() == "ctrl+c" || msg.String() == "q" {
		// 	return m, tea.Quit
		// }
	}
	return m, nil
}

func (m model) View() string {
	return m.message + "\n\nPress 'q' or Ctrl+C to quit."
}

func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
