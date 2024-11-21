package stringz

import (
	"fmt"
	"strings"
)

// AlignRight aligns the given string to the right, at the given width (minimum = 4), padding with dots.
func AlignRight(s string, width int) string {
	if width < 4 {
		width = 4
	}

	r := []rune(s)

	if len(r) > width {
		return fmt.Sprintf("...%v", string(r[len(r)-width+3:]))
	}

	if len(r) < width {
		return strings.Repeat(".", width-len(r)) + string(r)
	}

	return string(r)
}

// EnsureSuffix ensures that the string has the given (non-repeated) suffix.
func EnsureSuffix(s, suffix string) string {
	return strings.TrimSuffix(s, suffix) + suffix
}

// EnsurePrefix ensures that the string has the given (non-repeated) prefix.
func EnsurePrefix(s, prefix string) string {
	return prefix + strings.TrimPrefix(s, prefix)
}
