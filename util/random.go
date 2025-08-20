package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generate random int between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generate random string length of n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomOwner generate random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generate random amount of money in between 0 to 1000
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currenies := []string{USD, EUR, CAD}
	n := len(currenies)
	return currenies[rand.Intn(n)]
}
func RandomEmail() string {
	return fmt.Sprintf("%s@gmail.com", RandomString(6))
}
