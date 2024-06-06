package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

var (
	// imeiList   = []string{"864866057337932", "863051066174797", "864866059027275"}
	imeiList   = []string{"864866057337932"}
	deviceList = make([]device, 0)
)

type model struct {
	allowedImei []string
	lastInput   int
	sub         chan string
	lock        *sync.Mutex
	textInput   textinput.Model
	err         error
	logs        string
}

func initialModel() *model {
	ti := textinput.New()
	ti.Placeholder = "864866057337932"
	ti.Focus()
	ti.CharLimit = 15
	ti.Width = 30

	return &model{
		allowedImei: make([]string, 0),
		logs:        "",
		lastInput:   0,
		sub:         make(chan string, 10),
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
			m.lastInput++

		case tea.KeyCtrlC:
			return m, tea.Quit

		case tea.KeyCtrlR:
			if len(m.allowedImei) > 0 {
				m.allowedImei = m.allowedImei[:len(m.allowedImei)-1]
				if m.lastInput > 0 {
					m.lastInput--
				}
			}
		}

	case error:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(message)
	return m, cmd
}

func (m *model) View() string {
	s := ""

	s += "Enter your IMEI (15 digits): \n\n"

	s += fmt.Sprintf("Last Input: %d\n\n", m.lastInput)

	s += fmt.Sprintf("currently selected IMEIs:\n")
	s += fmt.Sprintf("%v", m.allowedImei) + "\n\n"

	if len(m.allowedImei) > 0 {
		s += "you will see:\n"

		s += strings.Repeat("=", 100) + "\n"
		s += m.logs
		select {
		case res := <-m.sub:
			m.logs += "\n" + res + "\n"
			// s += m.logs
			m.lastInput--
		default:
		}
		s += strings.Repeat("=", 100) + "\n"

		s += "\n"
	}

	s += m.textInput.View()
	s += "\n\n"
	s += "press ctrl+c to exit.\n"

	return s
}

type device struct {
	imei  string
	count int
}

func (m *model) logger(s string, d device) {
	for _, item := range m.allowedImei {
		if d.imei == item {
			m.sub <- s
			m.lastInput++
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
		// fmt.Println("allowed imeis 2: ", m.allowedImei, d.imei, fmt.Sprintf("%p", m))
	}
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Println("could not start program:", err)
		os.Exit(1)
	}
}
