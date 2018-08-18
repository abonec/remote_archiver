package main

import (
	"os"
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
