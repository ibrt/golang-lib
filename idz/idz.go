package idz

import (
	"github.com/google/uuid"
)

// MustNewRandomID generates a new random ID.
func MustNewRandomID() string {
	return uuid.Must(uuid.NewRandom()).String()
}

// IsValidID returns true if the given ID is valid.
func IsValidID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}
