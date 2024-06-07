package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	// imeiList   = []string{"864866057337932", "863051066174797", "864866059027275"}
	imeiList   = []string{"864866057337932"}
	deviceList = make([]device, 0)
)

type model struct {
	allowedImei []string
	lock        *sync.Mutex
	textInput   textinput.Model
	err         error
	maxLogLines int
	logs        []string
}

func initialModel() *model {
	ti := textinput.New()
	ti.Placeholder = "864866057337932"
	ti.Focus()
	ti.CharLimit = 15
	ti.Width = 30

	return &model{
		allowedImei: make([]string, 0),
		logs:        make([]string, 0),
		maxLogLines: 9,
		lock:        &sync.Mutex{},
		textInput:   ti,
		err:         nil,
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

	switch msg := message.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if len(m.textInput.Value()) != 15 {
				return m, nil
			}

			m.allowedImei = append(m.allowedImei, m.textInput.Value())
			m.textInput.SetValue("")
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyCtrlR:
			if len(m.allowedImei) > 0 {
				m.allowedImei = m.allowedImei[:len(m.allowedImei)-1]
			}
		}
	case tea.WindowSizeMsg:
		// fmt.Println(style.Render(fmt.Sprintf("%v", msg)))
	case error:
		m.err = msg
		return m, nil
	case cursor.BlinkMsg:
		if len(m.logs) > m.maxLogLines {
			m.logs = m.logs[1:]
		}
	}

	m.textInput, cmd = m.textInput.Update(message)
	return m, cmd
}

func (m *model) View() string {
	s := ""

	s += "Enter your IMEI (15 digits):\n"
	s += fmt.Sprintf("%v", m.allowedImei) + "\n\n"

	if len(m.allowedImei) > 0 && len(m.logs) > 0 {
		s += "LOGS:\n"
		s += logWin.Render(strings.Join(m.logs, "\n"))
		s += "\n"
	}

	s += m.textInput.View()
	s += "\n"

	exitStyle := lipgloss.NewStyle().
		Background(redPinkI).
		Foreground(whiteI)
	s += exitStyle.Render("press ctrl+c to exit")

	return mainWin.Render(s)
}

type device struct {
	imei  string
	count int
}

func (m *model) logger(s string, d device) {
	m.lock.Lock()
	defer m.lock.Unlock()
	for _, item := range m.allowedImei {
		if d.imei == item {
			m.logs = append(m.logs, s)
		}
	}
}

func (d device) produce(m *model) {
	// ticker := time.NewTicker(time.Duration(rand.Intn(9)+1) * time.Second)
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
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Println("could not start program:", err)
		os.Exit(1)
	}
}
