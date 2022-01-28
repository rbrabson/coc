package rest

import "fmt"

// ErrHttp is an error from an HTTP request
type ErrHttp struct {
	StatusCode int
	Status     string
}

// Error returns a formatted HTTP error
func (err ErrHttp) Error() string {
	return fmt.Sprintf("HTTP error: status=%d, reason=%s", err.StatusCode, err.Status)
}
