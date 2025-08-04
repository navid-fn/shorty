package utils

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func GeneratePseudoRandomString(length int) string {
	const letters = "abcdefghijklmnopqrstuvwxyz"

	// Create a new random generator seeded with the current time.
	// This is important!
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var sb strings.Builder
	for range length {
		// Pick a random character from the letters string
		sb.WriteByte(letters[r.Intn(len(letters))])
	}
	return sb.String()
}

func WriteJson(w http.ResponseWriter,status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}


func WriteError(w http.ResponseWriter, status int, msg string) {
	WriteJson(w, status, map[string]string{"error": msg})
}
