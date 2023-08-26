package helpers

import (
	"crypto/md5"
	"fmt"
)

func MD5Hash(s string) string {
	b := md5.Sum([]byte(s))
	return fmt.Sprintf("%x", b)
}
