package utils

import (
	"math/rand"
	"strings"
	"time"
)

func generatePseudoRandomString(length int) string {
	const letters = "abcdefghijklmnopqrstuvwxyz"
	
	// Create a new random generator seeded with the current time.
	// This is important!
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var sb strings.Builder
	for i := 0; i < length; i++ {
		// Pick a random character from the letters string
		sb.WriteByte(letters[r.Intn(len(letters))])
	}
	return sb.String()
}
