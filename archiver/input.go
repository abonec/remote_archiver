package archiver

import "io"

type Input interface {
	Reader() io.Reader
	Path() string
	Close()
}
