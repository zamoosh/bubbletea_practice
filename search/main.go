package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	boxWidth  = 50
	boxHeight = 10
)

type mode int

const (
	modeNormal mode = iota
	modeSearch
)

type inputModel struct {
	textInput textinput.Model
}

type displayModel struct {
	viewport viewport.Model
	content  string
}

type searchModel struct {
	textInput textinput.Model
	results   []int
	active    bool
}

type model struct {
	input   inputModel
	display displayModel
	search  searchModel
	mode    mode
}

func initialModel() *model {
	ti := textinput.New()
	ti.Placeholder = "Enter something..."
	ti.Focus()
	ti.Width = boxWidth - 4 // Adjust for padding/margin if needed

	vp := viewport.New(boxWidth, boxHeight)
	content := generateLongText()
	vp.SetContent(content)

	searchTI := textinput.New()
	searchTI.Placeholder = "Search..."
	searchTI.Width = boxWidth - 4

	return &model{
		input: inputModel{
			textInput: ti,
		},
		display: displayModel{
			viewport: vp,
			content:  content,
		},
		search: searchModel{
			textInput: searchTI,
			results:   []int{},
			active:    false,
		},
		mode: modeNormal,
	}
}

func (m *model) Init() tea.Cmd {
	return textinput.Blink
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.mode {
		case modeNormal:
			switch msg.String() {
			case "enter":
				m.display.viewport.SetContent(m.display.viewport.View() + "\n" + m.input.textInput.Value())
				m.input.textInput.SetValue("")
			case "/":
				m.mode = modeSearch
				m.search.active = true
				m.search.textInput.Focus()
			case "ctrl+c", "q":
				return m, tea.Quit
			}
		case modeSearch:
			switch msg.String() {
			case "enter":
				m.search.results = searchInContent(m.display.content, m.search.textInput.Value())
				m.search.active = false
				m.mode = modeNormal
				m.display.viewport.SetContent(highlightMatches(m.display.content, m.search.results))
				m.search.textInput.SetValue("")
			case "esc":
				m.search.active = false
				m.mode = modeNormal
			}
		}
	}

	var cmd tea.Cmd
	if m.mode == modeNormal {
		m.input.textInput, cmd = m.input.textInput.Update(msg)
	} else {
		m.search.textInput, cmd = m.search.textInput.Update(msg)
	}
	cmds = append(cmds, cmd)

	m.display.viewport, cmd = m.display.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m *model) View() string {
	inputView := m.input.textInput.View()
	if m.mode == modeSearch {
		inputView = m.search.textInput.View()
	}
	displayView := lipgloss.NewStyle().
		Width(boxWidth).
		Height(boxHeight).
		Border(lipgloss.RoundedBorder()).
		Padding(1, 2).
		Render(m.display.viewport.View())

	return inputView + "\n\n" + displayView + "\n\nPress 'q' or Ctrl+C to quit. Press '/' to search."
}

func generateLongText() string {
	lines := []string{}
	for i := 0; i < 100; i++ {
		lines = append(lines, fmt.Sprintf("This is line number %d", i+1))
	}
	return strings.Join(lines, "\n")
}

func searchInContent(content, pattern string) []int {
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringIndex(content, -1)
	positions := []int{}
	for _, match := range matches {
		positions = append(positions, match[0])
	}
	return positions
}

func highlightMatches(content string, positions []int) string {
	highlightStyle := lipgloss.NewStyle().Background(lipgloss.Color("yellow")).Foreground(lipgloss.Color("black"))
	for _, pos := range positions {
		content = content[:pos] + highlightStyle.Render(content[pos:pos+1]) + content[pos+1:]
	}
	return content
}

func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
