package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const number = "0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
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

func RandomOwner() string {
	return RandomString(8)
}

func RandomStoreName() string {
	return fmt.Sprintf("Store-%s", RandomString(4))
}

func RandomTableNumber() int64 {
	return RandomInt(0, 100)
}

func RandomTableName() string {
	return fmt.Sprintf("Table-%s", RandomString(2))
}

func RandomMenuName() string {
	return fmt.Sprintf("Menu-%s", RandomString(2))
}

func RandomPhone() string {
	var phone_number strings.Builder
	k := len(number)
	for i := 0; i < 10; i++ {
		c := number[rand.Intn(k)]
		phone_number.WriteByte(c)
	}
	return phone_number.String()
}
