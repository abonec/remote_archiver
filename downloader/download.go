package downloader

import (
	"github.com/abonec/file_downloader/archiver"
	"github.com/abonec/file_downloader/log"
	"os"
	"strconv"
	"sync"
	"net/http"
	"fmt"
)

func getParallel() int {
	env := os.Getenv("PARALLEL")
	if env == "" {
		return 2
	}
	size, err := strconv.Atoi(env)
	if log.Error(err) {
		os.Exit(1)
	}
	return size
}

func Download(inputQueue <-chan Input, verbose bool) <-chan archiver.Input {
	parallel := getParallel()
	fmt.Printf("parallel: %d\n", parallel)
	ch := make(chan archiver.Input, parallel)
	var wg sync.WaitGroup
	for i := 0; i < parallel; i++ {
		wg.Add(1)
		go func() {
			for input := range inputQueue {
				resp, err := http.Get(input.Url())
				if log.Error(err) {
					log.Warningf("Error while downloading %s", input.Url())
				}
				ch <- &Result{resp.Body, input.Path()}
			}
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	return ch
}
