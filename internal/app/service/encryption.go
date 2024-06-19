package service

import (
	"crypto/md5"
	"encoding/hex"
)

func HashText(text []byte) string {
	md5Hash := md5.Sum(text)
	hash := hex.EncodeToString(md5Hash[:])
	shortHash := hash[0:8]

	return shortHash
}
