package speed

import (
	"errors"
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

type SpeedTest struct {
	Ping     int64
	Download *TestResult
	Upload   *TestResult
}

func NewSpeedTest() *SpeedTest {
	test := &SpeedTest{}
	return test
}

func (s *SpeedTest) Start() (*SpeedTest, error) {

	if err := s.testDownload(); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *SpeedTest) testDownload() error {

	url := "https://proof.ovh.net/files/10Mb.dat"

	downloadStart := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		return errors.New(fmt.Sprintf("Err in Downloading: %v", err))
	}
	defer resp.Body.Close()

	duration := time.Since(downloadStart).Seconds()
	nBytes, err := io.Copy(io.Discard, resp.Body)
	if err != nil {
		return errors.New(fmt.Sprintf("Unable to count Bytes: %v", err))
	}

	mbps := float64(nBytes*8) / duration / 1_000_000
	buffer := float64(nBytes) / 1_000_000

	s.Download = &TestResult{
		Mbps:     mbps,
		Duration: duration,
		Buffer:   buffer,
	}

	return nil
}
