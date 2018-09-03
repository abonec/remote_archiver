package downloader

import "io"

type Result struct {
	reader io.Reader
	path   string
}

func (r *Result) Reader() io.Reader{
	return r.reader
}

func (r *Result) Path() string {
	return r.path
}
