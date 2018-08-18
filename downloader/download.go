package downloader

import (
	"net/http"
	"github.com/abonec/file_downloader/archiver"
	"github.com/abonec/file_downloader/log"
)

func Download(inputQueue <-chan Input) <-chan archiver.Input {
	ch := make(chan archiver.Input)
	go func() {
		for input := range inputQueue {
			resp, err := http.Get(input.Url())
			if log.Error(err) {
				log.Warningf("Error while downloading %s", input.Url())
			}
			ch <- &Result{resp.Body, input.Path()}
		}
		close(ch)
	}()
	return ch
}
