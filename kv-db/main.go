package main

import (
	"encoding/gob"
	"fmt"
	"hash/fnv"
	"os"

	lru "github.com/hashicorp/golang-lru"
)

type Page struct {
	Data map[string]string
}

type KVStore struct {
	Cache    *lru.Cache
	PageSize int
}

func NewKVStore(pageSize, cacheSize int) *KVStore {
	cache, _ := lru.New(cacheSize)
	return &KVStore{
		Cache:    cache,
		PageSize: pageSize,
	}
}

func (kv *KVStore) hashKey(key string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(key))
	return h.Sum32()
}

func (kv *KVStore) Get(key string) (string, bool) {
	hash := kv.hashKey(key)
	page, ok := kv.loadPage(hash)
	if !ok {
		return "", false
	}
	value := page.Data[key]
	return value, ok
}

func (kv *KVStore) Set(key string, value string) {
	hash := kv.hashKey(key)
	page, ok := kv.loadPage(hash)
	if !ok || len(page.Data) >= kv.PageSize {
		page = &Page{
			Data: make(map[string]string, kv.PageSize),
		}
	}
	page.Data[key] = value
	kv.savePage(hash, page)
}

func (kv *KVStore) loadPage(hash uint32) (*Page, bool) {
	// Try to get the page from the cache.
	if page, ok := kv.Cache.Get(hash); ok {
		return page.(*Page), true
	}

	// If it's not in the cache, try to load it from disk.
	file, err := os.Open(fmt.Sprint(hash))
	if err != nil {
		return nil, false
	}
	defer file.Close()

	page := &Page{}
	decoder := gob.NewDecoder(file)
	if err := decoder.Decode(page); err != nil {
		return nil, false
	}

	// Add the loaded page to the cache.
	kv.Cache.Add(hash, page)

	return page, true
}

func (kv *KVStore) savePage(hash uint32, page *Page) {
	// Save the page to the cache.
	kv.Cache.Add(hash, page)

	// Also save the page to disk.
	file, _ := os.Create(fmt.Sprint(hash))
	defer file.Close()

	encoder := gob.NewEncoder(file)
	encoder.Encode(page)
}

func main() {
	// Usage
	kv := NewKVStore(100, 1000)
	kv.Set("key", "value")
	value, _ := kv.Get("key")
	println(value)
}
