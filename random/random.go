package random

import (
	"crypto/rand"
	"fmt"

	"github.com/ibrt/golang-lib/errors"
)

// Handy alphabets.
var (
	AlphaNum      = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	LowerAlphaNum = []rune("abcdefghijklmnopqrstuvwxyz0123456789")
)

// Default ID length.
const (
	IDLen = 16
)

// GetSecureBytes returns a byte slice of cryptographically secure random numbers of the given "length"
func GetSecureBytes(length int) ([]byte, error) {
	buf := make([]byte, length)
	if _, err := rand.Read(buf); err != nil {
		return nil, errors.Wrap(err, errors.Skip())
	}
	return buf, nil
}

// MustGetSecureBytes is like GetSecureBytes but panics on error.
func MustGetSecureBytes(length int) []byte {
	buf, err := GetSecureBytes(length)
	errors.MaybeMustWrap(err, errors.Skip())
	return buf
}

// GetHex gets a hex-encoded random string of the given length (in digits).
func GetHex(length int) (string, error) {
	buf, err := GetSecureBytes(length / 2)
	if err != nil {
		return "", errors.Wrap(err, errors.Skip())
	}
	return fmt.Sprintf("%x", buf), nil
}

// MustGetHex is like GetHex but panics on error.
func MustGetHex(length int) string {
	s, err := GetHex(length)
	errors.MaybeMustWrap(err, errors.Skip())
	return s
}

// GetAlphabet gets a random string of the given length (in runes) composed of the given alphabet.
// The alphabet must be between 2 and 256 unique runes long.
func GetAlphabet(length int, alphabet []rune) (string, error) {
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
		if _, err := rand.Read(readBuf); err != nil {
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
func MustGetAlphabet(length int, alphabet []rune) string {
	s, err := GetAlphabet(length, alphabet)
	errors.MaybeMustWrap(err, errors.Skip())
	return s
}

// GetID calls GetAlphabet with sane default settings: 16-long, URL-safe lowercase, about 82 bits of entropy.
func GetID() (string, error) {
	return GetAlphabet(IDLen, LowerAlphaNum)
}

// MustGetID is like GetID but panics on error.
func MustGetID() string {
	s, err := GetAlphabet(IDLen, LowerAlphaNum)
	errors.MaybeMustWrap(err, errors.Skip())
	return s
}
