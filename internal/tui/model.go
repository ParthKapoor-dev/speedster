package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	status string
}

func InitModel() *Model {
	return &Model{
		status: "Welcome to Speedster CLI",
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}
