package archiver

import (
	"archive/zip"
	"fmt"
	"github.com/abonec/file_downloader/config"
	"github.com/abonec/file_downloader/log"
	"io"
	"os"
	"path/filepath"
)

const baseDir = "export"

func Archive(inputQueue <-chan Input, cfg config.Config) (io.Reader, <-chan int64) {
	pr, pw := io.Pipe()

	arch := zip.NewWriter(pw)
	ch := make(chan int64)
	go func() {
		defer pw.Close()
		defer arch.Close()
		i := 0
		for input := range inputQueue {
			header := &zip.FileHeader{
				Name:   filepath.Join(baseDir, input.Path()),
				Method: zip.Store,
			}
			writer, err := arch.CreateHeader(header)
			if log.Error(err) {
				os.Exit(1)
			}
			n, err := io.Copy(writer, input.Reader())
			ch <- n
			i++
			if cfg.Verbose() {
				fmt.Printf("%d files archived\r", i)
			}
		}
		fmt.Println()
	}()
	return pr, ch
}
