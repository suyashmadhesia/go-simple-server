package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabets = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// generates a random number between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min)
}

// function to generate random string
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabets)
	for i := 0; i < n; i++ {
		c := alphabets[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// function to generate random username
func RandomOwner() string {
	return RandomString(6)
}

// function Random currency generates
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "INR"}
	return currencies[rand.Intn(len(currencies))]
}

// function to generate random amount
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}
