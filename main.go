package main

import (
	"flag"
	"fmt"
	"github.com/abonec/file_downloader/archiver"
	"github.com/abonec/file_downloader/downloader"
	"github.com/abonec/file_downloader/parser"
	"github.com/abonec/file_downloader/uploader"
	"os"
	"runtime/trace"
	"os/signal"
	"syscall"
)

// TODO: cancel upload if there is no input files
func main() {
	//InitLogger(os.Stdout)
	uploadKey := flag.String("upload_key", "", "Key for upload")
	verbose := flag.Bool("verbose", false, "Verbose mode")
	startTrace := flag.Bool("trace", false, "Collect trace information")
	flag.Parse()
	if *uploadKey == "" {
		fmt.Println("need to specify upload_key; see -h")
		os.Exit(1)
	}

	if *startTrace {
		file, err := os.Create("trace.out")
		if err != nil {
			fmt.Println(err)
		}
		trace.Start(file)
		defer trace.Stop()
	}

	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-sigs
		trace.Stop()
		fmt.Println("Interrupting...")
		os.Exit(0)
	}()

	inputQueue := parser.Parse(os.Stdin, *verbose)
	downloadQueue := downloader.Download(inputQueue, *verbose)
	reader := archiver.Archive(downloadQueue, *verbose)
	uploader.Upload(reader, *uploadKey, *verbose)
}
