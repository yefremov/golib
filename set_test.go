package golib_test

import (
	"fmt"
	"testing"

	"github.com/yefremov/golib"
)

func TestSet(t *testing.T) {
	s := golib.NewSet("foo", "bar", "baz")
	if s.Size() != 3 {
		t.Error("Expected size 3")
	}
	fmt.Println(s)
	s.Add()
	if s.Size() != 3 {
		t.Error("Expected size 3")
	}

	s.Add(1, 2, 3)
	if s.Size() != 6 {
		t.Error("Expected size 6")
	}

	s.Remove()
	if s.Size() != 6 {
		t.Error("Expected size 6")
	}

	s.Remove("foo", 2, "bux")
	if s.Size() != 4 {
		t.Error("Expected size 4")
	}

	if s.Pop() == nil {
		t.Error("Expeted value")
	}

	if s.Has("baz") == false {
		t.Error("Expeted to have baz")
	}

	if s.Has() {
		t.Error("Expeted false")
	}

	s.Clear()
	if s.Pop() != nil {
		t.Error("Expeted nil")
	}

	if s.Has("baz") == true {
		t.Error("Expeted to not have baz")
	}

	s.Add(1, 2, 3)
	s.Each(func(item interface{}) bool {
		return true
	})

	s.List()
}
