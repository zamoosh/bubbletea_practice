package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	boxWidth  = 50
	boxHeight = 10
)

type inputModel struct {
	textInput textinput.Model
}

type displayModel struct {
	viewport viewport.Model
}

type model struct {
	input   inputModel
	display displayModel
}

func initialModel() *model {
	ti := textinput.New()
	ti.Placeholder = "Enter something..."
	ti.Focus()
	ti.Width = boxWidth - 4 // Adjust for padding/margin if needed

	vp := viewport.New(boxWidth, boxHeight)

	return &model{
		input: inputModel{
			textInput: ti,
		},
		display: displayModel{
			viewport: vp,
		},
	}
}

func (m *model) Init() tea.Cmd {
	return textinput.Blink
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			m.display.viewport.SetContent(m.display.viewport.View() + "\n" + m.input.textInput.Value())
			m.input.textInput.SetValue("")
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.input.textInput, cmd = m.input.textInput.Update(msg)
	cmds = append(cmds, cmd)

	m.display.viewport, cmd = m.display.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m *model) View() string {
	inputView := m.input.textInput.View()
	displayView := lipgloss.NewStyle().
		Width(boxWidth).
		Height(boxHeight).
		Border(lipgloss.RoundedBorder()).
		Padding(1, 2).
		Render(m.display.viewport.View())

	return inputView + "\n\n" + displayView + "\n\nPress 'q' or Ctrl+C to quit."
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
