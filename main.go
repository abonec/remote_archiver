package main

import (
	"fmt"
	"github.com/abonec/file_downloader/archiver"
	"github.com/abonec/file_downloader/config"
	"github.com/abonec/file_downloader/downloader"
	"github.com/abonec/file_downloader/notifier"
	"github.com/abonec/file_downloader/parser"
	"github.com/abonec/file_downloader/tracing"
	"github.com/abonec/file_downloader/uploader"
	"os"
	"runtime/trace"
)

// TODO: cancel upload if there is no input files
func main() {
	cfg, err := config.Init()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if tracing.Start(cfg) != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer trace.Stop()

	inputQueue := parser.Parse(os.Stdin)
	downloadQueue, downloadSizeChan := downloader.Download(inputQueue)
	reader, archivedSizeChan := archiver.Archive(downloadQueue, cfg)

	go notifier.RunNotifier(downloadSizeChan, archivedSizeChan)
	uploader.Upload(reader, cfg)
}
