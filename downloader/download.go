package downloader

import (
	"github.com/abonec/file_downloader/archiver"
	"github.com/abonec/file_downloader/log"
	"os"
	"strconv"
	"sync"
	"net/http"
	"fmt"
	"time"
	"io"
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

type DownloadBody struct {
	io.ReadCloser
	wait chan interface{}
}

func newDownloadBody(body io.ReadCloser) *DownloadBody {
	return &DownloadBody{ReadCloser: body, wait: make(chan interface{})}
}

func (db *DownloadBody) Close() error {
	err := db.ReadCloser.Close()
	close(db.wait)
	return err
}

func (db *DownloadBody) WaitClose() {
	<-db.wait
}

func Download(inputQueue <-chan Input, verbose bool) <-chan archiver.Input {
	parallel := getParallel()
	fmt.Printf("parallel: %d\n", parallel)
	ch := make(chan archiver.Input, parallel)
	var wg sync.WaitGroup
	for i := 0; i < parallel; i++ {
		wg.Add(1)
		go func() {
			transport := &http.Transport{
			}
			client := &http.Client{
				Timeout:   60 * time.Second,
				Transport: transport,
			}
			for input := range inputQueue {
				resp, err := retryableDownload(client, input.Url())
				if log.Error(err) {
					log.Warningf("Error while downloading %s", input.Url())
				}
				body := newDownloadBody(resp.Body)
				ch <- &Result{body, input.Path()}
				body.WaitClose()
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

const RETRIES = 5

func retryableDownload(client *http.Client, url string) (*http.Response, error) {
	var err error
	for tries := 0; tries < RETRIES; tries++ {
		resp, err := client.Get(url)
		if err != nil {
			fmt.Println("got exception", err)
			continue
		}
		return resp, nil
	}
	return nil, err
}
