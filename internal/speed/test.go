package speed

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type TestResult struct {
	Mbps     float64
	Duration float64
	Buffer   float64
}

type ProgressMsg struct {
	BytesDownloaded int64
	Percent         float64
	Elapsed         float64
}

type SpeedTest struct {
	Ping            int64
	Download        *TestResult
	Upload          *TestResult
	progressHandler func(progress ProgressMsg)
}

func NewSpeedTest() *SpeedTest {
	test := &SpeedTest{}
	return test
}

func (s *SpeedTest) Start(onProgress func(msg ProgressMsg)) (*SpeedTest, error) {

	s.progressHandler = onProgress

	if err := s.testDownload(); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *SpeedTest) testDownload() error {
	url := "https://proof.ovh.net/files/10Mb.dat"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Cache-Control", "no-cache")

	client := &http.Client{
		// Timeout: 30 * time.Second, // Consider uncommenting this for a robust solution
		Transport: &http.Transport{
			DisableKeepAlives: true,
		},
	}

	downloadStart := time.Now()
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error in downloading: %v", err)
	}
	defer resp.Body.Close()

	const totalSize = 10 * 1024 * 1024 // 10 MB

	progressReader := &ProgressReader{
		Reader:    resp.Body,
		TotalSize: totalSize,
		Start:     time.Now(),
		OnProgress: func(downloaded int64, percent float64, elapsed float64) {
			s.progressHandler(ProgressMsg{
				BytesDownloaded: downloaded,
				Percent:         percent,
				Elapsed:         elapsed,
			})
		},
	}

	nBytes, err := io.Copy(io.Discard, progressReader)

	if err != nil {
		return fmt.Errorf("unable to count bytes: %v", err)
	}
	duration := time.Since(downloadStart).Seconds()

	s.Download = &TestResult{
		Mbps:     float64(nBytes*8) / duration / 1_000_000,
		Duration: duration,
		Buffer:   float64(nBytes) / 1_000_000,
	}

	return nil
}
