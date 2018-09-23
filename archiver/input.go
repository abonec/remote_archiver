package archiver

import "io"

type Input interface {
	Path() string
	io.ReadCloser
}
