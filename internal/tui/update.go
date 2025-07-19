package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/parthkapoor-dev/speedster/internal/speed"
	"github.com/parthkapoor-dev/speedster/internal/tui/status"
)

type SpeedTestResultMsg struct {
	Result *speed.SpeedTest
}

type SpeedTestErrorMsg struct {
	Err error
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		// start speed test
		case "s":
			m.status = status.Start
			m.message = "Starting speed test"

			return m, StartSpeedTestCmd(m.Program)

		}

	case speed.ProgressMsg:
		m.message = fmt.Sprintf("Downloading... %.1f%% complete (%.2f MB / 10 MB) in %.1fs", msg.Percent, float64(msg.BytesDownloaded)/1_000_000, msg.Elapsed)
		return m, nil

	case SpeedTestErrorMsg:
		m.status = status.Err
		m.message = error.Error(msg.Err)
		return m, nil

	case SpeedTestResultMsg:
		m.status = status.Success
		m.message = "Successfully Tested"
		m.downloadResult = msg.Result.Download
		return m, nil
	}

	return m, nil
}

func StartSpeedTestCmd(p *tea.Program) tea.Cmd {
	return func() tea.Msg {
		test := speed.NewSpeedTest()
		result, err := test.Start(
			func(progMsg speed.ProgressMsg) {
				p.Send(progMsg)
			},
		)
		if err != nil {
			return SpeedTestErrorMsg{Err: err}
		}
		return SpeedTestResultMsg{Result: result}
	}
}
