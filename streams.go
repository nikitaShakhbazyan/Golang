package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	t := time.Now()
	rand.Seed(t.UnixNano())

	parseURL("http://example.com")
	parseURL("http://youtube.com")

	fmt.Println("parsing is finished",time.Since(t).Seconds())

}

func parseURL(url string) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5; i++ {
		latency := rand.Intn(500) + 500

		time.Sleep(time.Duration(latency) * time.Millisecond)

		fmt.Printf("Parsing <%s> - Step %d - Latency: %dms\n", url, i+1, latency)
	}
}