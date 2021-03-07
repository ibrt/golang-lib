package random

import (
	"context"
	"crypto/rand"
	"fmt"
	"io"

	"github.com/ibrt/golang-lib/errors"
	"github.com/ibrt/golang-lib/inject"
)

type contextKey int

const (
	randomContextKey contextKey = iota
)

// Initializer is a Random initializer.
func Initializer(_ context.Context) (inject.Injector, inject.Releaser, error) {
	return inject.SingletonInjectorFactory(randomContextKey, NewRandom(rand.Reader)), nil, nil
}

// SingletonInjectorFactory always injects the given Random.
func SingletonInjectorFactory(rnd *Random) inject.Injector {
	return inject.SingletonInjectorFactory(randomContextKey, rnd)
}

// Get returns the Random, or nil if not found.
func Get(ctx context.Context) *Random {
	if rnd, ok := ctx.Value(randomContextKey).(*Random); ok {
		return rnd
	}
	return nil
}

// MustGet returns the Random, panics if not found.
func MustGet(ctx context.Context) *Random {
	clk := Get(ctx)
	errors.Assertf(clk != nil, "random unexpectedly nil", errors.Skip())
	return clk
}

// Handy alphabets.
var (
	AlphaNum      = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	LowerAlphaNum = []rune("abcdefghijklmnopqrstuvwxyz0123456789")
)

// Default ID length.
const (
	IDLen = 16
)

// Random provides random generation utilities.
type Random struct {
	reader io.Reader
}

// NewRandom initializes a new Random.
func NewRandom(reader io.Reader) *Random {
	return &Random{
		reader: reader,
	}
}

// Read implements the io.Reader interface, proxying requests to the underlying reader.
func (r *Random) Read(buf []byte) (int, error) {
	n, err := r.reader.Read(buf)
	return n, errors.MaybeWrap(err, errors.Skip())
}

// GetSecureBytes returns a byte slice of cryptographically secure random numbers of the given "length"
func (r *Random) GetSecureBytes(length int) ([]byte, error) {
	buf := make([]byte, length)
	if _, err := r.reader.Read(buf); err != nil {
		return nil, errors.Wrap(err, errors.Skip())
	}
	return buf, nil
}

// MustGetSecureBytes is like GetSecureBytes but panics on error.
func (r *Random) MustGetSecureBytes(length int) []byte {
	buf, err := r.GetSecureBytes(length)
	errors.MaybeMustWrap(err, errors.Skip())
	return buf
}

// GetHex gets a hex-encoded random string of the given length (in digits).
func (r *Random) GetHex(length int) (string, error) {
	buf, err := r.GetSecureBytes(length / 2)
	if err != nil {
		return "", errors.Wrap(err, errors.Skip())
	}
	return fmt.Sprintf("%x", buf), nil
}

// MustGetHex is like GetHex but panics on error.
func (r *Random) MustGetHex(length int) string {
	s, err := r.GetHex(length)
	errors.MaybeMustWrap(err, errors.Skip())
	return s
}

// GetAlphabet gets a random string of the given length (in runes) composed of the given alphabet.
// The alphabet must be between 2 and 256 unique runes long.
func (r *Random) GetAlphabet(length int, alphabet []rune) (string, error) {
	if length == 0 {
		return "", nil
	}
	if len(alphabet) < 2 || len(alphabet) > 256 {
		return "", errors.Errorf("invalid alphabet length: %v", errors.A(len(alphabet)), errors.Skip())
	}

	maxR := byte(255 - (256 % len(alphabet))) // used to avoid modulo bias
	buf := make([]rune, 0, length)
	readBuf := make([]byte, length+(length/4))

	for {
		if _, err := r.reader.Read(readBuf); err != nil {
			return "", errors.Wrap(err, errors.Skip())
		}
		for _, r := range readBuf {
			if r <= maxR {
				if buf = append(buf, alphabet[int(r)%len(alphabet)]); len(buf) == length {
					return string(buf), nil
				}
			}
		}
	}
}

// MustGetAlphabet is like GetAlphabet but panics on error.
func (r *Random) MustGetAlphabet(length int, alphabet []rune) string {
	s, err := r.GetAlphabet(length, alphabet)
	errors.MaybeMustWrap(err, errors.Skip())
	return s
}

// GetID calls GetAlphabet with sane default settings: 16-long, URL-safe lowercase, about 82 bits of entropy.
func (r *Random) GetID() (string, error) {
	return r.GetAlphabet(IDLen, LowerAlphaNum)
}

// MustGetID is like GetID but panics on error.
func (r *Random) MustGetID() string {
	s, err := r.GetAlphabet(IDLen, LowerAlphaNum)
	errors.MaybeMustWrap(err, errors.Skip())
	return s
}
