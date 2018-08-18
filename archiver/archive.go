package archiver

import (
	"io"
	"os"
	"github.com/abonec/file_downloader/log"
	"archive/zip"
	"path/filepath"
	)

const baseDir = "export"

func Archive(inputQueue <-chan Input) io.ReadSeeker {
	zipFile, err := os.Create("test_arch.zip")
	if log.Error(err) {
		os.Exit(1)
	}

	arch := zip.NewWriter(zipFile)
	//go func() {
		defer arch.Close()
		for input := range inputQueue {
			header := &zip.FileHeader{
				Name:   filepath.Join(baseDir, input.Path()),
				Method: zip.Store,
			}
			writer, err := arch.CreateHeader(header)
			if log.Error(err){
				os.Exit(1)
			}
			_, err = io.Copy(writer, input.Reader())
		}

	//}()

	return zipFile
}
