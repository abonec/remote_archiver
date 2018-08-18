package main

import (
	"io/ioutil"
	"os"
)

func setDownloadDir(dir *string) {
	if *dir == "" {
		tmp, err := ioutil.TempDir("", "downloader")
		if logError(err) {
			os.Exit(1)
		}
		*dir = tmp
	}
}
