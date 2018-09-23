package downloader

import (
	"bytes"
	"sync"
)

type Result struct {
	buffer *bytes.Buffer
	path   string
	pool   *sync.Pool
}

func (r *Result) Path() string {
	return r.path
}
func (r *Result) Read(p []byte) (n int, err error) {
	return r.buffer.Read(p)
}

func (r *Result) Close() error {
	r.pool.Put(r.buffer)
	return nil
}
