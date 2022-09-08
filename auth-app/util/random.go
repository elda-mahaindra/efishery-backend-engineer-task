package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const alphanumeric = "1234567890"
const chars = alphabet + alphanumeric

func init() {
	rand.Seed(time.Now().UnixNano())
}

// primary

func RandomAlphanumeric(n int) string {
	var sb strings.Builder
	k := len(alphanumeric)

	for i := 0; i < n; i++ {
		c := alphanumeric[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomBool() bool {
	return rand.Intn(2) == 1
}

func RandomChars(n int) string {
	var sb strings.Builder
	k := len(chars)

	for i := 0; i < n; i++ {
		c := chars[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomIntFromSet(a ...int) int {
	n := len(a)
	if n == 0 {
		return 0
	}

	return a[rand.Intn(n)]
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomStringFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}

	return a[rand.Intn(n)]
}

// secondary

func RandomName() string {
	return RandomString(9)
}

func RandomPassword() string {
	return RandomChars(4)
}

func RandomPhone() string {
	return RandomAlphanumeric(13)
}
