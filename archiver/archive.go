package archiver

import (
	"io"
	"os"
	"github.com/abonec/file_downloader/log"
	"archive/zip"
	"path/filepath"
	"fmt"
	)

const baseDir = "export"

func Archive(inputQueue <-chan Input, verbose bool) io.Reader {
	pr, pw := io.Pipe()

	arch := zip.NewWriter(pw)
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
			_, err = io.Copy(writer, input.Reader())
			i++
			if verbose {
				fmt.Printf("%d files archived\r", i)
			}
			input.Reader().Close()
		}
		fmt.Println()
	}()
	return pr
}
