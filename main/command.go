package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const url = "https://charm.sh/"

type model struct {
	status     int
	err        error
	helpMsg    string
	showHelp   bool
	processing bool
}

func initialModel() model {
	return model{
		status:     0,
		err:        nil,
		helpMsg:    "This program only sends a request to the server and shows status code\n\n",
		showHelp:   false,
		processing: false,
	}
}

func makeRequest() tea.Cmd {
	return func() tea.Msg {
		c := &http.Client{Timeout: 10 * time.Second}
		res, err := c.Get(url)
		if err != nil {
			return err
		}
		return res.StatusCode
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			m.processing = true
			return m, makeRequest()

		case "ctrl+c", "q":
			return m, tea.Quit

		case "ctrl+h":
			m.showHelp = true
			return m, nil

		case "esc":
			m.showHelp = false
			m.processing = false
			return m, nil
		}

	case int:
		m.status = msg
		return m, tea.Quit

	case error:
		m.err = msg
		return m, tea.Quit
	}
	return m, nil
}

func (m model) View() string {
	s := ""

	if m.processing {
		s = fmt.Sprintf("Checking %s ... ", url)

		if m.err != nil {
			return fmt.Sprintf("\nWe had some trouble: %v\n\n", m.err)
		}

		if m.status > 0 {
			s += fmt.Sprintf("%d %s!", m.status, http.StatusText(m.status))
		}
	} else if m.showHelp {
		s = "Just press the Enter buddy!\n\n"
		s += "Press esc go to main menu\n"
		return s
	} else {
		s = "This program only sends a request to 'https://charm.sh/' and shows status code.\n\n"
		s += "Press enter to run\n"
		s += "Press q to quit\n"
		s += "Press esc go to main menu\n"
		s += "Press ctrl+h to show help message\n"

		if m.showHelp {
			return m.helpMsg
		}
	}

	return s + "\n"
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an err: %v", err)
		os.Exit(1)
	}
}
