package tui

import (
	"animatic-v2/Structures"
	"errors"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	AnimeNameColumnTitle = "Anime Name"
	AnimeURLColumnTitle  = "Anime URL"
)

type SelectAnimeModel struct {
	table             table.Model
	selectedURL       string
	selectedAnimeName string
}

func (m *SelectAnimeModel) Init() tea.Cmd {
	return nil
}

func (m *SelectAnimeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			tea.Quit()
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			m.selectedAnimeName = m.table.SelectedRow()[0]
			m.selectedURL = m.table.SelectedRow()[1]
			return m, tea.Quit
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m *SelectAnimeModel) View() string {
	return m.table.View()
}

func SelectAnimes(animes []Structures.Anime) (string, string, error) {
	if len(animes) == 0 {
		return "", "", errors.New("Anime List is empty")
	}

	columns := []table.Column{
		{Title: AnimeNameColumnTitle, Width: 50},
		{Title: AnimeURLColumnTitle, Width: 180},
	}

	rows := make([]table.Row, 0, len(animes))

	for _, anime := range animes {
		rows = append(rows, table.Row{anime.Name, anime.Url})
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(len(animes)),
	)

	s := table.DefaultStyles()

	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)

	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)

	t.SetStyles(s)

	m := &SelectAnimeModel{table: t}
	p := tea.NewProgram(m)
	if err := p.Start(); err != nil {
		return "", "", err
	}

	return m.selectedAnimeName, m.selectedURL, nil
}
