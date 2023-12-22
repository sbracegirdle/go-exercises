package main

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	cache, err := NewLRUCache(2)

	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	cache.Put(1, 1)
	cache.Put(2, 2)
	if val, err := cache.Get(1); val != 1 || err != nil {
		t.Errorf("Expected 1, got %d", val)
	}

	cache.Put(3, 3) // This will cause key 2 to be evicted
	if val, err := cache.Get(2); val != -1 || err == nil {
		t.Errorf("Expected -1, got %d", val)
	}

	cache.Put(4, 4) // This will cause key 1 to be evicted
	if val, err := cache.Get(1); val != -1 || err == nil {
		t.Errorf("Expected -1, got %d", val)
	}
	if val, err := cache.Get(3); val != 3 || err != nil {
		t.Errorf("Expected 3, got %d", val)
	}
	if val, err := cache.Get(4); val != 4 || err != nil {
		t.Errorf("Expected 4, got %d", val)
	}
}

func TestLRUCache_UpdateExistingKey(t *testing.T) {
	cache, err := NewLRUCache(2)

	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(1, 10) // Update the value of key 1

	if val, err := cache.Get(1); val != 10 || err != nil {
		t.Errorf("Expected 10, got %d", val)
	}
}

func TestLRUCache_GetNonExistingKey(t *testing.T) {
	cache, err := NewLRUCache(2)

	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	cache.Put(1, 1)
	cache.Put(2, 2)

	if val, err := cache.Get(3); val != -1 || err == nil {
		t.Errorf("Expected -1, got %d", val)
	}
}

func TestLRUCache_CapacityZero(t *testing.T) {
	cache, err := NewLRUCache(0)

	if cache != nil {
		t.Errorf("Expected nil, got %v", cache)
	}

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
