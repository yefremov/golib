package golib

import "sync"

var (
	keyExists = struct{}{}
)

type Set interface {
	Add(items ...interface{})
	Remove(items ...interface{})
	Pop() interface{}
	Has(items ...interface{}) bool
	Size() int
	Clear()
	Each(f func(interface{}) bool)
	List() []interface{}
}

type set struct {
	items map[interface{}]struct{}
	mutex sync.RWMutex
}

func NewSet(items ...interface{}) Set {
	s := &set{mutex: sync.RWMutex{}}
	s.items = make(map[interface{}]struct{})
	s.Add(items...)
	return s
}

func (s *set) Add(items ...interface{}) {
	if len(items) == 0 {
		return
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for _, item := range items {
		s.items[item] = keyExists
	}
}
func (s *set) Remove(items ...interface{}) {
	if len(items) == 0 {
		return
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for _, item := range items {
		delete(s.items, item)
	}
}
func (s *set) Pop() interface{} {
	s.mutex.RLock()
	for item := range s.items {
		s.mutex.RUnlock()
		s.mutex.Lock()
		delete(s.items, item)
		s.mutex.Unlock()
		return item
	}
	s.mutex.RUnlock()
	return nil
}

func (s *set) Has(items ...interface{}) bool {
	// assume checked for empty item, which not exist
	if len(items) == 0 {
		return false
	}

	s.mutex.RLock()
	defer s.mutex.RUnlock()

	has := true
	for _, item := range items {
		if _, has = s.items[item]; !has {
			break
		}
	}
	return has
}

func (s *set) Size() int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return len(s.items)
}

func (s *set) Clear() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.items = make(map[interface{}]struct{})
}

func (s *set) Each(f func(item interface{}) bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	for item := range s.items {
		if !f(item) {
			break
		}
	}
}

func (s *set) List() []interface{} {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	list := make([]interface{}, 0, len(s.items))

	for item := range s.items {
		list = append(list, item)
	}

	return list
}
