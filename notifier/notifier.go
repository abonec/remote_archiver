package notifier

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"sync/atomic"
	"time"
)

func RunNotifier(downloadSizeChan, uploadSizeChan <-chan int64) {
	var downloadSize int64
	var uploadSize int64
	startTime := time.Now()

	go func() {
		for {
			select {
			case down := <-downloadSizeChan:
				atomic.AddInt64(&downloadSize, down)
			case up := <-uploadSizeChan:
				atomic.AddInt64(&uploadSize, up)
			}
		}
	}()

	for {
		time.Sleep(1 * time.Second)
		duration := int64(time.Since(startTime).Seconds())
		downSpeed := atomic.LoadInt64(&downloadSize) / duration
		uploadSpeed := atomic.LoadInt64(&uploadSize) / duration
		fmt.Printf("\033[2K\rdownload: %s/s, upload: %s/s", humanize.Bytes(uint64(downSpeed)), humanize.Bytes(uint64(uploadSpeed)))
	}

}
