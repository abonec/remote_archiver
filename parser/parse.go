package parser

import (
	"io"
	"github.com/abonec/file_downloader/downloader"
	"strings"
	"bufio"
)

// Parse parses input for path and urls and return channel with the results.
// Input should be as one line with relative path for file and url where file
// should be downloaded. Path and url should be joined with triple semicolon ;;;
// i.e.: path/to.jpg;;;http://example.com/from.jpg
func Parse(input io.Reader) <-chan downloader.Input {
	queue := make(chan downloader.Input)

	go func() {
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			line := scanner.Text()
			data := strings.Split(line, ";;;")
			if len(data) != 2 {
				continue
			}
			queue <- &Result{data[0], data[1]}
		}
		close(queue)
	}()

	return queue
}
