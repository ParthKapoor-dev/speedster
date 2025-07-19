package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/parthkapoor-dev/speedster/internal/tui"
)

func main() {
	m := tui.InitModel()
	p := tea.NewProgram(m)

	m.Program = p

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
