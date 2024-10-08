package lfucache

import "fmt"

type Set[T comparable] struct {
	keys map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		keys: make(map[T]struct{}),
	}
}

func (set *Set[T]) Add(value T) {
	set.keys[value] = struct{}{}
}

func (set *Set[T]) Remove(value T) {
	delete(set.keys, value)
}

func (set *Set[T]) PopAny() (ans T) {
	for key := range set.keys {
		set.Remove(key)
		return key
	}

	return ans
}

func (set *Set[T]) IsEmpty() bool {
	return len(set.keys) == 0
}

func (set *Set[T]) String() string {
	return fmt.Sprintf("%v", set.keys)
}
