package golib_test

import (
	"testing"

	"github.com/yefremov/golib"
)

func TestCache(t *testing.T) {
	var value interface{}
	var result interface{}
	var ok bool

	cache := golib.NewCache(3)

	value = nil

	result = cache.Get(0)
	if result != value {
		t.Errorf("Expected value %v got %v", value, result)
	}

	value = "foo"

	ok = cache.Put(0, value)
	if !ok {
		t.Errorf("Expected successful put operation got %v", ok)
	}

	result = cache.Get(0)
	if result != value {
		t.Errorf("Expected value %v got %v", value, result)
	}

	value = "bar"

	ok = cache.Put(0, value)
	if !ok {
		t.Errorf("Expected successful put operation got %v", ok)
	}

	ok = cache.Put(0, nil)
	if ok {
		t.Errorf("Expected unsuccessful put operation got %v", ok)
	}

	cache.Put(0, "a")
	cache.Put(1, "b")
	cache.Put(2, "c")
	cache.Put(3, "d")

	result = cache.Get(0)
	if result != nil {
		t.Errorf("Expected value %v got %v", nil, result)
	}
}
