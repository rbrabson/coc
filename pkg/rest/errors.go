package rest

import "fmt"

// ErrHttp is an error from an HTTP request
type ErrHttp struct {
	URL        string
	StatusCode int
	Status     string
}

// Error returns a formatted HTTP error
func (err ErrHttp) Error() string {
	return fmt.Sprintf("HTTP error: url=%s, status=%d, reason=%s", err.URL, err.StatusCode, err.Status)
}
