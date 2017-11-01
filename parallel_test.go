package golib_test

import (
	"sync"
	"sync/atomic"
	"testing"

	"github.com/yefremov/golib"
)

func TestParallel(t *testing.T) {
	var wg sync.WaitGroup
	var ops uint64 = 0
	var n int = 100
	wg.Add(n)
	golib.Parallel(n, func() {
		defer wg.Done()
		atomic.AddUint64(&ops, 1)
	})
	wg.Wait()
	if atomic.LoadUint64(&ops) != uint64(n) {
		t.Errorf("Given %v not euqals executed %v", n, &ops)
	}
}
