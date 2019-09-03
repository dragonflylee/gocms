package util

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5 生成32位MD5
func MD5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
