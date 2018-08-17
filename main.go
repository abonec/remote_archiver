package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Bool("download", true, "download mode")
	uploadMode := flag.Bool("upload", false, "Upload mode")
	uploadFile := flag.String("upload_file", "", "File for upload")
	uploadKey := flag.String("upload_key", "", "Key for upload")
	flag.Parse()

	if *uploadMode {
		if *uploadFile == "" || *uploadKey == "" {
			fmt.Println("For uploading keys upload_file and upload_key should be specified. See -h")
			os.Exit(1)
		}
		runUpload(*uploadFile, *uploadKey)
	} else {
		runDownload()
	}
}
