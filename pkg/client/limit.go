package client

import (
	"errors"
	"io"
)

// ErrResponseTooLarge is the error of a response whose body is larger than the limit set by
// Client.WithMaxResponseBytes.
var ErrResponseTooLarge = errors.New("the response body is larger than the limit")

// limitedReader reads at most left bytes of r. Unlike io.LimitedReader it does not end quietly at the limit: a body
// that goes on is an error, so a cut body is never taken for the whole one.
type limitedReader struct {
	r    io.ReadCloser
	left int64
}

func (l *limitedReader) Read(p []byte) (int, error) {
	if l.left <= 0 {
		// The limit is reached. The body is complete only if there is nothing more.
		var probe [1]byte
		n, err := l.r.Read(probe[:])
		if n > 0 {
			return 0, ErrResponseTooLarge
		}
		return 0, err
	}
	if int64(len(p)) > l.left {
		p = p[:l.left]
	}
	n, err := l.r.Read(p)
	l.left -= int64(n)
	return n, err
}

func (l *limitedReader) Close() error { return l.r.Close() }
