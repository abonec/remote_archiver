package main

import (
	"os"
	"flag"
	"fmt"
)

func main() {
	uploadKey := flag.String("upload_key", "", "Key for upload")
	downloadDir := flag.String("download_dir", "", "Dir for download")
	setDownloadDir(downloadDir)
	flag.Parse()
	if *uploadKey == "" {
		fmt.Printf("need to specify upload_key; see -h")
		os.Exit(1)
	}
	runDownload(*downloadDir)

	zip, err := archive(*downloadDir)
	if logError(err){
		os.Exit(1)
	}
	err = os.RemoveAll(*downloadDir)
	if logError(err) {
		os.Exit(1)
	}
	runUpload(zip, *uploadKey)
	err = os.Remove(zip)
	if logError(err) {
		os.Exit(1)
	}
}
