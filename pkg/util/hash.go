package util

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
)

// MD5 生成32位MD5
func MD5(a ...string) string {
	h := md5.New()
	for _, v := range a {
		h.Write([]byte(v))
	}
	return hex.EncodeToString(h.Sum(nil))
}

func RandString(n int) string {
	b := make([]byte, n)
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(b)
}
