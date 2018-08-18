package downloader

type Input interface {
	Path() string
	Url() string
}
