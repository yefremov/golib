package golib

import "sync"

func Parallel(n int, fn func()) {
	var wg sync.WaitGroup
	wg.Add(n)
	defer wg.Wait()

	for i := 0; i < n; i++ {
		go func() {
			fn()
			wg.Done()
		}()
	}
}
