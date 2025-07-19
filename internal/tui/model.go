package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/parthkapoor-dev/speedster/internal/speed"
	"github.com/parthkapoor-dev/speedster/internal/tui/status"
)

type Model struct {
	Program        *tea.Program
	message        string
	status         status.Status
	downloadResult *speed.TestResult
}

func InitModel() *Model {
	return &Model{
		status:  status.Ok,
		message: "Hello World! press `s` to start the speed-test",
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}
