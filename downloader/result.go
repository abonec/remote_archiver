package downloader

import (
	"bytes"
	"io"
	"sync"
)

type Result struct {
	buffer *bytes.Buffer
	path   string
	pool   *sync.Pool
}

func (r *Result) Reader() io.Reader {
	return r.buffer
}

func (r *Result) Path() string {
	return r.path
}

func (r *Result) Close() {
	r.pool.Put(r.buffer)
}
