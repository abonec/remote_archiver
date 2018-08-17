package main

import (
	"os"
	"path/filepath"
	"strconv"
)

//func getEnv(key string) string {
//	if env := os.Getenv(key); env != "" {
//		return env
//	}
//	return ""
//}

func getParallel() int {
	env := os.Getenv("PARALLEL")
	if env == "" {
		return 2
	}
	size, err := strconv.Atoi(env)
	if logError(err) {
		os.Exit(1)
	}
	return size
}

func getDownloadDir() string {
	var dir string
	if env := os.Getenv("DOWNLOAD_DIR"); env != "" {
		dir = env
	} else {
		dir = "./downloads"
	}
	abs, err := filepath.Abs(dir)
	if logError(err) {
		os.Exit(1)
	}
	return abs
}
