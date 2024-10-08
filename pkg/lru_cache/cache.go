package lrucache

//Implemented LRU earlier, so decided why not to add.

type IList[K comparable, V any] interface {
	PushFront(node *Node[K, V])
	Remove(node *Node[K, V])
	Back() *Node[K, V]
	MoveToFront(node *Node[K, V])
}

type LRU[K comparable, V any] struct {
	cap   int
	list  IList[K, V]
	cache map[K]*Node[K, V]
}

func NewLRU[K comparable, V any](cap int) *LRU[K, V] {
	return &LRU[K, V]{
		cap:   cap,
		list:  NewList[K, V](),
		cache: make(map[K]*Node[K, V]),
	}
}

func (lru *LRU[K, V]) Get(key K) (value V, found bool) {
	if node, ok := lru.cache[key]; ok {
		lru.list.MoveToFront(node)
		return node.Value, true
	}

	return value, false
}

func (lru *LRU[K, V]) Put(key K, value V) {
	if node, ok := lru.cache[key]; ok {
		node.Value = value
		lru.list.MoveToFront(node)
	}

	if len(lru.cache) == lru.cap {
		back := lru.list.Back()
		if back != nil {
			lru.list.Remove(back)
			delete(lru.cache, back.Key)
		}
	}

	node := NewNode(key, value)
	lru.list.PushFront(node)
	lru.cache[key] = node
}
