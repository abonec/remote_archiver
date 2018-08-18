package parser

type Result struct {
	path string
	url  string
}

// Path returns relative path to for file where it
// should be saved
func (r *Result) Path() string {
	return r.path
}

// Url returns url where file should be downloaded
func (r *Result) Url() string {
	return r.url
}
