package crypto_tool

import (
	"crypto/md5"
	"encoding/hex"
)

// GetMD5 return input to MD5 hash
func GetMD5(input string) string {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}
