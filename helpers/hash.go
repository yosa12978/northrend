package helpers

import (
	"crypto/md5"
	"fmt"
	"math/rand"
)

func MD5Hash(s string) string {
	b := md5.Sum([]byte(s))
	return fmt.Sprintf("%x", b)
}

func NewSalt() string {
	salt := make([]byte, 16)
	rand.Read(salt)
	return fmt.Sprintf("%x", salt)
}
