package controller

import (
	lru "github.com/hashicorp/golang-lru"
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
