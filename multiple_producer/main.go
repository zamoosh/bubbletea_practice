package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	imeiList   = []string{"864866057337932", "863051066174797", "864866059027275"}
	deviceList = make([]device, 0)
)

type model struct {
	allowedImei []string
	lock        *sync.Mutex
	textInput   textinput.Model
	err         error
	maxLogLines int
	logs        []string
	display     *logDisplay
}

type logDisplay struct {
	viewport       viewport.Model
	scrollToBottom bool
}

type device struct {
	imei  string
	count int
}

func initialModel() *model {
	ti := textinput.New()
	ti.Placeholder = "12345678912345"
	ti.Focus()
	ti.CharLimit = 15
	ti.Width = 15

	vp := viewport.New(100, 18)

	return &model{
		allowedImei: make([]string, 0),
		logs:        make([]string, 0, 20),
		maxLogLines: 200,
		lock:        &sync.Mutex{},
		textInput:   ti,
		err:         nil,
		display: &logDisplay{
			viewport:       vp,
			scrollToBottom: true,
		},
	}
}

func (m *model) Init() tea.Cmd {
	for _, item := range imeiList {
		d := device{imei: item, count: 500}
		deviceList = append(deviceList, d)

		go d.produce(m)
	}

	return textinput.Blink
}

func (m *model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var commands []tea.Cmd

	switch msg := message.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "alt+b":
			m.display.scrollToBottom = !m.display.scrollToBottom
		}

		switch msg.Type {
		case tea.KeyEnter:
			if len(m.textInput.Value()) != 15 {
				return m, nil
			}

			m.allowedImei = append(m.allowedImei, m.textInput.Value())
			m.textInput.SetValue("")
			// m.textInput.Blur()
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyCtrlR:
			if len(m.allowedImei) > 0 {
				m.allowedImei = m.allowedImei[:len(m.allowedImei)-1]
				m.logs = []string{}
				m.display.viewport.SetContent("")
			}
		}
	case error:
		m.err = msg
		return m, nil
	case cursor.BlinkMsg:
		if len(m.logs) > m.maxLogLines {
			m.logs = m.logs[1:]
		}
		m.display.viewport.SetContent(strings.Join(m.logs, "\n"))
		if m.display.scrollToBottom {
			m.display.viewport.GotoBottom()
		}
	case tea.MouseMsg:
		switch message.(type) {
		}
	}

	m.textInput, cmd = m.textInput.Update(message)
	commands = append(commands, cmd)

	m.display.viewport, cmd = m.display.viewport.Update(message)
	commands = append(commands, cmd)

	return m, tea.Batch(commands...)
}

func (m *model) View() string {
	s := ""

	s += "Enter your IMEI (15 digits): "
	s += imeiListStyle.Render(fmt.Sprintf("%v", m.allowedImei))
	s += "\n"

	s += "Scroll to bottom: "
	if m.display.scrollToBottom {
		s += lipgloss.NewStyle().Foreground(greenI).Render(fmt.Sprintf("%v", m.display.scrollToBottom))
	} else {
		s += lipgloss.NewStyle().Foreground(redPinkI).Render(fmt.Sprintf("%v", m.display.scrollToBottom))
	}
	s += "\n"

	if len(m.allowedImei) > 0 && len(m.logs) > 0 {
		s += "LOGS:\n"
		s += logWin.Render(m.display.viewport.View())
		s += "\n"
	} else {
		s += strings.Repeat("\n", 18)
	}

	s += m.textInput.View()
	s += "\n"

	exitStyle := lipgloss.NewStyle().
		Background(redPinkI).
		Foreground(whiteI).
		PaddingLeft(1).
		PaddingRight(1)

	s += exitStyle.Render("press ctrl+c to exit")

	return mainWin.Render(s)
}

func (m *model) logger(s string, d device) {
	m.lock.Lock()
	defer m.lock.Unlock()
	for _, item := range m.allowedImei {
		if strings.Contains(s, item) {
			m.logs = append(m.logs, s)
		}
	}
}

func (d device) produce(m *model) {
	// ticker := time.NewTicker(time.Duration(rand.Intn(4)+1) * time.Second)
	ticker := time.NewTicker(1 * time.Second)

	for d.count > 0 {
		now := time.Now().UTC().Format("20060102150405")
		msg := fmt.Sprintf("$G1,%s,%v,1,34.570601,50.810350,1042,0.00,,12.84,,,,00890,69,58,15,13,*95", d.imei, now)

		<-ticker.C
		d.count--
		m.logger(msg, d)
	}
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("could not start program:", err)
		os.Exit(1)
	}
}
