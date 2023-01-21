package cache

import (
	"crypto/md5"
	"encoding/hex"
)

func GenerateHash(str string) string {
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}
