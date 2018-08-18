package main

import (
	"github.com/abonec/file_downloader/downloader"
	"github.com/abonec/file_downloader/parser"
	"os"
	"github.com/abonec/file_downloader/archiver"
)

func main() {
	//InitLogger(os.Stdout)
	//uploadKey := flag.String("upload_key", "", "Key for upload")
	//downloadDir := flag.String("download_dir", "", "Dir for download")
	//verbose := flag.Bool("verbose", false, "show progress of downloading")
	//setDownloadDir(downloadDir)
	//flag.Parse()
	//if *uploadKey == "" {
	//	fmt.Printf("need to specify upload_key; see -h")
	//	os.Exit(1)
	//}
	//runDownload(*downloadDir, *verbose)
	//
	//zip, err := archive(*downloadDir)
	//if logError(err){
	//	os.Exit(1)
	//}
	//err = os.RemoveAll(*downloadDir)
	//if logError(err) {
	//	os.Exit(1)
	//}
	//runUpload(zip, *uploadKey)
	//err = os.Remove(zip)
	//if logError(err) {
	//	os.Exit(1)
	//}

	inputQueue := parser.Parse(os.Stdin)
	downloadQueue := downloader.Download(inputQueue)
	archiver.Archive(downloadQueue)
}
