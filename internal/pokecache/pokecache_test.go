package pokecache

import (
	"testing"
	"time"
)

func TestCache_Add(t *testing.T) {
	cache := NewCache(time.Millisecond * 10)
	cache.Add("test", []byte("test"))
	if _, ok := cache.cache["test"]; !ok {
		t.Errorf("expected cache to have key %s", "test")
	}
}

func TestCache_Get(t *testing.T) {
	cache := NewCache(time.Millisecond * 10)
	cache.Add("test", []byte("test"))
	if _, ok := cache.Get("test"); !ok {
		t.Errorf("expected cache to have key %s", "test")
	}
}

func TestCache_Reap(t *testing.T) {
	cache := NewCache(time.Millisecond * 10)
	cache.Add("test", []byte("test"))
	time.Sleep(time.Millisecond * 10 * 2)

	if _, ok := cache.cache["test"]; ok {
		t.Errorf("expected cache to not have key %s", "test")
	}

}
