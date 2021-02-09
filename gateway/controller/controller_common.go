package controller

import (
	"encoding/base64"
	"encoding/hex"

	"github.com/Eric-GreenComb/contrib/crypto"
	"github.com/hashicorp/golang-lru"
)

var lruCache *lru.ARCCache

func init() {
	lruCache, _ = lru.NewARC(8192)
}

// ClearLruCache ClearLruCache
func ClearLruCache() {
	lruCache.Purge()
}

// RemoveLruCache RemoveLruCache
func RemoveLruCache(key interface{}) {
	lruCache.Remove(key)
}

// DesDecryptHex DesDecryptHex
func DesDecryptHex(encHex, key string) (string, error) {
	_dec, err := hex.DecodeString(encHex)
	if err != nil {
		return "", err
	}

	_orig, err := crypto.DesDecrypt(_dec, []byte(key))
	if err != nil {
		return "", err
	}
	return string(_orig), nil
}

// DesDecryptBase64 DesDecryptBase64
func DesDecryptBase64(encBase64, key string) (string, error) {
	_dec, err := base64.StdEncoding.DecodeString(encBase64)
	if err != nil {
		return "", err
	}
	_orig, err := crypto.DesDecrypt(_dec, []byte(key))
	if err != nil {
		return "", err
	}
	return string(_orig[:]), nil
}
