package main

import (
	"fmt"
	"github.com/abonec/file_downloader/archiver"
	"github.com/abonec/file_downloader/config"
	"github.com/abonec/file_downloader/downloader"
	"github.com/abonec/file_downloader/parser"
	"github.com/abonec/file_downloader/tracing"
		"os"
	"runtime/trace"
	"github.com/abonec/file_downloader/uploader"
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
	downloadQueue := downloader.Download(inputQueue)
	reader := archiver.Archive(downloadQueue, cfg)
	uploader.Upload(reader, cfg)
}
