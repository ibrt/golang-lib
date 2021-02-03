package random_test

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"testing"

	"github.com/ibrt/golang-lib/random"
	"github.com/stretchr/testify/require"
)

func TestGetSecureBytes(t *testing.T) {
	orig := rand.Reader
	defer func() {
		rand.Reader = orig
	}()

	readBuf := make([]byte, 16)
	for i := 0; i < len(readBuf); i++ {
		readBuf[i] = byte(i)
	}
	rand.Reader = bytes.NewReader(readBuf)

	require.Empty(t, random.MustGetSecureBytes(0))
	buf := random.MustGetSecureBytes(16)
	require.Len(t, buf, 16)
	require.Equal(t, readBuf[:16], buf)

	buf, err := random.GetSecureBytes(16)
	require.EqualError(t, err, "EOF")
	require.Nil(t, buf)
	require.Panics(t, func() { _ = random.MustGetSecureBytes(16) })
}

func TestGetHex(t *testing.T) {
	orig := rand.Reader
	defer func() {
		rand.Reader = orig
	}()

	readBuf := make([]byte, 8)
	for i := 0; i < len(readBuf); i++ {
		readBuf[i] = byte(i)
	}
	rand.Reader = bytes.NewReader(readBuf)

	require.Empty(t, random.MustGetHex(0))
	s := random.MustGetHex(16)
	require.Len(t, s, 16)
	require.Equal(t, fmt.Sprintf("%x", readBuf[:8]), s)

	s, err := random.GetHex(16)
	require.EqualError(t, err, "EOF")
	require.Empty(t, s)
	require.Panics(t, func() { _ = random.MustGetHex(16) })
}

func TestGetAlphabet(t *testing.T) {
	orig := rand.Reader
	defer func() {
		rand.Reader = orig
	}()

	readBuf := make([]byte, len(random.AlphaNum)+len(random.AlphaNum)/4)
	for i := 0; i < len(readBuf); i++ {
		readBuf[i] = byte(i)
	}
	rand.Reader = bytes.NewReader(readBuf)

	require.Empty(t, random.MustGetAlphabet(0, random.AlphaNum))
	s := random.MustGetAlphabet(len(random.AlphaNum), random.AlphaNum)
	require.Len(t, s, len(random.AlphaNum))
	require.Equal(t, string(random.AlphaNum), s)

	s, err := random.GetAlphabet(16, random.AlphaNum)
	require.EqualError(t, err, "EOF")
	require.Empty(t, s)
	require.Panics(t, func() { _ = random.MustGetAlphabet(16, random.AlphaNum) })
	require.Panics(t, func() { random.MustGetAlphabet(16, []rune{}) })
}

func TestGetID(t *testing.T) {
	orig := rand.Reader
	defer func() {
		rand.Reader = orig
	}()

	readBuf := make([]byte, random.IDLen+random.IDLen/4)
	for i := 0; i < len(readBuf); i++ {
		readBuf[i] = byte(i)
	}
	rand.Reader = bytes.NewReader(readBuf)

	s := random.MustGetID()
	require.Len(t, s, random.IDLen)
	require.Equal(t, string(random.LowerAlphaNum)[:random.IDLen], s)

	s, err := random.GetID()
	require.EqualError(t, err, "EOF")
	require.Empty(t, s)
	require.Panics(t, func() { _ = random.MustGetID() })
}
