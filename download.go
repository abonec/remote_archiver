package main

import (
	"bufio"
	"fmt"
	"github.com/dustin/go-humanize"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
)

type DownloadTask struct {
	root        string
	path        string
	url         string
	resultQueue chan *DownloadResult
}

func NewTask(root, path, url string, result chan *DownloadResult) *DownloadTask {
	return &DownloadTask{root, path, url, result}
}

func (dt *DownloadTask) Run() {
	absRoot, err := filepath.Abs(dt.root)
	if logError(err) {
		return
	}
	fullPath := path.Join(absRoot, dt.path)
	dirPath := filepath.Dir(fullPath)
	os.MkdirAll(dirPath, os.ModePerm)

	out, err := os.Create(fullPath)
	if logError(err) {
		return
	}
	defer out.Close()

	resp, err := http.Get(dt.url)
	if logError(err) {
		return
	}
	defer resp.Body.Close()

	size, err := io.Copy(out, resp.Body)

	if logError(err) {
		return
	}
	dt.resultQueue <- &DownloadResult{Url: dt.url, Path: dt.path, Size: size}

}

type DownloadResult struct {
	Url  string
	Path string
	Size int64
}

var parallel = getParallel()

func runDownload(downloadDir string, verbose bool) {
	Trace.Println("Downloading started")
	scanner := bufio.NewScanner(os.Stdin)

	queue := make(chan *DownloadTask)
	result := make(chan *DownloadResult)
	fmt.Printf("Downloading to %s\n", downloadDir)
	fmt.Printf("Parallel downloading: %d\n", parallel)
	var wg sync.WaitGroup
	for i := 0; i < parallel; i++ {
		wg.Add(1)
		go func() {
			for {
				task, more := <-queue
				if !more {
					wg.Done()
					return
				}
				task.Run()
			}
		}()
	}

	tasks := make([]*DownloadTask, 0)
	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Split(line, ";;;")
		if len(data) != 2 {
			fmt.Printf("Given string %s; shuld be relative path and url join by triple semicolon (path;;;url)\n", line)
			continue
		}
		tasks = append(tasks, NewTask(downloadDir, strings.TrimSpace(data[0]), strings.TrimSpace(data[1]), result))
	}
	fileSize := len(tasks)
	go func() {
		for _, task := range tasks {
			queue <- task
		}
		close(queue)
	}()
	go func() {
		wg.Wait()
		close(result)
		fmt.Println()
		Trace.Println("Downloading finished")
	}()

	var size int64
	var downloaded int
	for {
		result, more := <-result
		if !more {
			break
		}
		size += result.Size
		downloaded += 1
		if verbose {
			fmt.Printf("\r%d of %d, Total amout: %s; Downloaded: %s; %s                                    ",
				downloaded, fileSize, humanize.Bytes(uint64(size)), result.Url, humanize.Bytes(uint64(result.Size)))
		}
	}
}
