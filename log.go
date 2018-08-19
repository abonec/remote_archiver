package main

import (
	"fmt"
	"log"
	"io"
)

var (
	Trace *log.Logger
)

func InitLogger(traceHandle io.Writer) {
	Trace = log.New(traceHandle, "", log.Ltime)
}

func logError(err error) bool {
	if err != nil {
		fmt.Println(err)
		return true
	} else {
		return false
	}
}
