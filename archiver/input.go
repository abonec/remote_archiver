package archiver

import "io"

type Input interface {
	Reader() io.ReadCloser
	Path() string
}
