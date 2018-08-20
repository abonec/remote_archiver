package main

import (
	"flag"
	"fmt"
	"github.com/abonec/file_downloader/archiver"
	"github.com/abonec/file_downloader/downloader"
	"github.com/abonec/file_downloader/parser"
	"github.com/abonec/file_downloader/uploader"
	"os"
)

func main() {
	//InitLogger(os.Stdout)
	uploadKey := flag.String("upload_key", "", "Key for upload")
	verbose := flag.Bool("verbose", false, "Verbose mode")
	flag.Parse()
	if *uploadKey == "" {
		fmt.Println("need to specify upload_key; see -h")
		os.Exit(1)
	}

	inputQueue := parser.Parse(os.Stdin, *verbose)
	downloadQueue := downloader.Download(inputQueue, *verbose)
	reader := archiver.Archive(downloadQueue, *verbose)
	uploader.Upload(reader, *uploadKey, *verbose)
}
