package tui

import "fmt"

func (m Model) View() string {
	return fmt.Sprintf("%s\n\nPress q to quit.", m.status)
}
