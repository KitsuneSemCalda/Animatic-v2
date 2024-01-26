package tui

import (
	message "animatic-v2/Message"

	tea "github.com/charmbracelet/bubbletea"

	"time"
)

type searchLabelModel struct {
	choice string
	err    error
	done   bool
}

func (m *searchLabelModel) Init() tea.Cmd {
	return nil
}

func (m *searchLabelModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			m.done = true
			return m, tea.Quit
		case tea.KeyBackspace:
			if len(m.choice) > 0 {
				m.choice = m.choice[:len(m.choice)-1]
			}
		default:
			if msg.String() != "" {
				m.choice += msg.String()
			}
		}
	}
	return m, nil
}

func (m *searchLabelModel) View() string {
	return "Enter the name of the anime to be downloaded: " + m.choice
}

func GetAnimeName() string {
	m := &searchLabelModel{}
	p := tea.NewProgram(m)
	errChan := make(chan error)
	go func() {
		errChan <- p.Start()
	}()
	select {
	case err := <-errChan:
		if err != nil {
			message.ErrorMessage(err.Error())
			return ""
		}
	case <-time.After(time.Second * 10): // Timeout after 10 seconds
		p.Send(tea.KeyMsg{Type: tea.KeyCtrlC}) // Send Ctrl+C to the program
	}
	if m.done {
		return m.choice
	}
	message.ErrorMessage("Could not assert model to searchLabelModel")
	return ""
}
