package random

import (
	"crypto/rand"
	"fmt"

	"github.com/ibrt/golang-lib/errors"
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

// GetHex gets a hex-encoded random string of the given length (in characters).
func GetHex(length int) (string, error) {
	buf, err := GetSecureBytes(length / 2)
	if err != nil {
		return "", errors.Wrap(err, errors.Skip())
	}
	return fmt.Sprintf("%v", buf[:length]), nil
}

// MustGetHex is like GetHex but panics on error.
func MustGetHex(length int) string {
	s, err := GetHex(length)
	errors.MaybeMustWrap(err)
	return s
}
