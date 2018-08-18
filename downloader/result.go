package downloader

import "io"

type Result struct {
	reader io.ReadCloser
	path   string
}

func (r *Result) Reader() io.ReadCloser {
	return r.reader
}

func (r *Result) Path() string {
	return r.path
}
