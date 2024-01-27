package tui

import (
	message "animatic-v2/Message"
	structure "animatic-v2/Structures"
	"os"
	"os/signal"
	"syscall"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type selectAnimeModel struct {
	cursor int
	animes []structure.Anime
	choice int
	err    error
}

func (m *selectAnimeModel) Init() tea.Cmd {
	return tea.Batch(tea.ClearScreen)
}

func (m *selectAnimeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyUp:
			if m.cursor > 0 {
				m.cursor--
			}
		case tea.KeyDown:
			if m.cursor < len(m.animes)-1 {
				m.cursor++
			}
		case tea.KeyEnter:
			m.choice = m.cursor
			return m, tea.Quit
		case tea.KeyCtrlC:
			os.Exit(0)
		}
	}
	return m, nil
}

func (m *selectAnimeModel) View() string {
	cursorStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF00")).Bold(true)
	listStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF"))
	s := listStyle.Render("Select the anime:\n")
	for i, anime := range m.animes {
		cursor := " "
		if m.cursor == i {
			cursor = cursorStyle.Render("▶")
		}
		s += cursor + " " + listStyle.Render(anime.Name) + "\n"
	}
	return s
}

func SelectAnimes(animes []structure.Anime) int {
	m := &selectAnimeModel{animes: animes} // Use um ponteiro para o modelo
	p := tea.NewProgram(m)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		tea.ClearScreen()
		tea.ShowCursor()
		os.Exit(0)
	}()

	if err := p.Start(); err != nil {
		message.ErrorMessage(err.Error())
		return -1
	}

	return m.choice
}
