package util

import (
	"math/rand"
	"strings"
)

const alphabets = "abcdefghijklmnopqrstuwxyz"

func RandomInt(max int64) int64 {
	return rand.Int63n(max)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabets)

	for i := 0; i < n; i++ {
		c := alphabets[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// function to return a random owner name all the time its invoked.
func RandomOwner() string {
	return RandomString(10)
}

// function to return a random money all the time its invoked.
func RandomMoney() int64 {
	return RandomInt(10000)
}

// function to return random currency
func RandomCurrency() string {
	currencies := []string{"INR", "USD", "EUR", "CAD"}
	n := len(currencies)

	return currencies[rand.Intn(n)]
}
