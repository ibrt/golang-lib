package ioz

import (
	"io"
	"sync/atomic"
)

var (
	_ io.Reader = (*CountingReader)(nil)
)

// CountingReader implements a io.Reader that counts bytes.
type CountingReader struct {
	r       io.Reader
	counter *atomic.Int64
}

// NewCountingReader initializes a new CountingReader.
func NewCountingReader(r io.Reader) *CountingReader {
	return &CountingReader{
		r:       r,
		counter: &atomic.Int64{},
	}
}

// Read implements the io.Reader interface.
func (c *CountingReader) Read(p []byte) (int, error) {
	n, err := c.r.Read(p)
	c.counter.Add(int64(n))
	return n, err
}

// Count returns the number of bytes read.
func (c *CountingReader) Count() int64 {
	return c.counter.Load()
}
