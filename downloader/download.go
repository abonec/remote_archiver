package downloader

import (
	"bytes"
	"fmt"
	"github.com/abonec/file_downloader/archiver"
	"github.com/abonec/file_downloader/log"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
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
	io.Reader
	wait chan interface{}
}

func (db *DownloadBody) Close() error {
	//err := db.ReadCloser.Close()
	close(db.wait)
	return nil
}

func (db *DownloadBody) WaitClose() {
	<-db.wait
}

func Download(inputQueue <-chan Input) (<-chan archiver.Input, <-chan int64) {
	pool := &sync.Pool{New: func() interface{} {
		//ar := make([]byte, 0, 10000)
		return new(bytes.Buffer)
	}}
	parallel := getParallel()
	fmt.Printf("parallel: %d\n", parallel)
	ch := make(chan archiver.Input, parallel)
	bytesDownloadedChan := make(chan int64)
	var wg sync.WaitGroup
	for i := 0; i < parallel; i++ {
		wg.Add(1)
		go func() {
			//transport := &http.Transport{
			//}
			client := &http.Client{
				Timeout: 60 * time.Second,
				//Transport: transport,
			}
			for input := range inputQueue {
				buffer, bytesDownloaded, err := retryableDownload(client, input.Url(), pool)
				if log.Error(err) {
					log.Warningf("Error while downloading %s", input.Url())
				}
				ch <- &Result{buffer: buffer, path: input.Path(), pool: pool}
				bytesDownloadedChan <- bytesDownloaded
			}
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	return ch, bytesDownloadedChan
}

const RETRIES = 5

func retryableDownload(client *http.Client, url string, pool *sync.Pool) (*bytes.Buffer, int64, error) {
	var err error
	for tries := 0; tries < RETRIES; tries++ {
		resp, err := client.Get(url)
		if err != nil {
			continue
		}
		buf := pool.Get().(*bytes.Buffer)
		buf.Reset()
		n, err := io.Copy(buf, resp.Body)
		resp.Body.Close()
		if err != nil {
			continue
		}
		return buf, n, nil
	}
	return nil, 0, err
}
