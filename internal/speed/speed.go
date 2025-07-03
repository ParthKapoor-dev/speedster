package speed

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type SpeedTest struct {
	ping     int64
	download int64
	upload   int64
}

func NewSpeedTest() *SpeedTest {
	test := &SpeedTest{}
	return test
}

func (s *SpeedTest) StartTest() {

	url := "https://proof.ovh.net/files/10Mb.dat"

	downloadStart := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		log.Println("Err in Downloading", err)
		return
	}
	defer resp.Body.Close()

	duration := time.Since(downloadStart).Seconds()
	nBytes, err := io.Copy(io.Discard, resp.Body)
	if err != nil {
		fmt.Println("Unable to count Bytes", err)
	}

	mbps := float64(nBytes*8) / duration / 1_000_000

	fmt.Printf("Downloaded %.2f MB in %.2f seconds\n", float64(nBytes)/1_000_000, duration)
	fmt.Printf("Download speed: %.2f Mbps\n", mbps)
}
