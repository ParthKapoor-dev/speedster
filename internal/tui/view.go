package tui

import (
	"fmt"

	"github.com/parthkapoor-dev/speedster/internal/speed"
)

func (m Model) View() string {

	if m.downloadResult != nil {
		return getDownloadResult(m.downloadResult)
	}

	return fmt.Sprintf("%s\n\nPress q to quit.", m.message)
}

func getDownloadResult(result *speed.TestResult) string {

	return fmt.Sprintf(
		"Downloaded %.2f MB in %.2f seconds\nDownload speed: %.2f Mbps\n",
		result.Buffer, result.Duration, result.Mbps)

}
