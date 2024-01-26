package tui

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func TestSearchLabelModelUpdate(t *testing.T) {
	m := &searchLabelModel{}
	msg := tea.KeyMsg{Type: tea.KeyEnter}
	model, _ := m.Update(msg)
	if model.(*searchLabelModel).done != true {
		t.Errorf("Expected done to be true, got %v", model.(*searchLabelModel).done)
	}
}

func TestSearchLabelModelView(t *testing.T) {
	m := &searchLabelModel{choice: "Naruto"}
	expected := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#00FF00")).
		Bold(true).
		Border(lipgloss.RoundedBorder()).
		Padding(1).
		Width(50).
		Render("Enter the name of the anime to be downloaded: Naruto")
	if m.View() != expected {
		t.Errorf("Expected %v, got %v", expected, m.View())
	}
}

func TestUpdateWithBackspace(t *testing.T) {
	m := &searchLabelModel{choice: "Naruto"}
	msg := tea.KeyMsg{Type: tea.KeyBackspace}
	model, _ := m.Update(msg)
	if model.(*searchLabelModel).choice != "Narut" {
		t.Errorf("Expected choice to be 'Narut', got %v", model.(*searchLabelModel).choice)
	}
}

func TestUpdateWithCharacter(t *testing.T) {
	m := &searchLabelModel{choice: "Naruto"}
	msg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("s")}
	model, _ := m.Update(msg)
	if model.(*searchLabelModel).choice != "Narutos" {
		t.Errorf("Expected choice to be 'Narutos', got %v", model.(*searchLabelModel).choice)
	}
}

func TestInit(t *testing.T) {
	m := &searchLabelModel{}
	cmd := m.Init()
	if cmd == nil {
		t.Errorf("Expected cmd to not be nil")
	}
}
