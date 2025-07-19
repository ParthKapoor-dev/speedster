package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/parthkapoor-dev/speedster/internal/tui"
)

func main() {

	if _, err := tea.NewProgram(tui.InitModel()).Run(); err != nil {
		log.Fatal(err)
	}

	// result, err := speed.NewSpeedTest().Start()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("Downloaded %.2f MB in %.2f seconds\n", result.Download.Buffer, result.Download.Duration)
	// fmt.Printf("Download speed: %.2f Mbps\n", result.Download.Mbps)
}
