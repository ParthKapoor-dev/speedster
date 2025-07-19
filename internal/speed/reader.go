package speed

import (
	"io"
	"time"
)

type ProgressReader struct {
	Reader       io.Reader
	TotalSize    int64
	BytesRead    int64
	Start        time.Time
	OnProgress   func(downloaded int64, percent float64, elapsed float64)
	lastReported time.Time
}

func (pr *ProgressReader) Read(p []byte) (int, error) {
	n, err := pr.Reader.Read(p)
	pr.BytesRead += int64(n)

	now := time.Now()
	if now.Sub(pr.lastReported) > 100*time.Millisecond || err != nil {
		pr.lastReported = now
		elapsed := now.Sub(pr.Start).Seconds()
		percent := (float64(pr.BytesRead) / float64(pr.TotalSize)) * 100
		pr.OnProgress(pr.BytesRead, percent, elapsed)
	}

	return n, err
}
