package lrucache

type Node[K comparable, V any] struct {
	Key   K
	Value V
	Next  *Node[K, V]
	Prev  *Node[K, V]
}

type List[K comparable, V any] struct {
	Head *Node[K, V]
	Tail *Node[K, V]
}

func NewNode[K comparable, V any](key K, value V) *Node[K, V] {
	return &Node[K, V]{
		Key:   key,
		Value: value,
		Next:  nil,
		Prev:  nil,
	}
}

func NewList[K comparable, V any]() *List[K, V] {
	var (
		key   K
		value V
	)

	list := &List[K, V]{
		Head: NewNode[K, V](key, value),
		Tail: NewNode[K, V](key, value),
	}

	list.Head.Next = list.Tail
	list.Tail.Prev = list.Head
	return list
}

func (list *List[K, V]) PushFront(node *Node[K, V]) {
	node.Next = list.Head.Next
	node.Prev = list.Head
	list.Head.Next.Prev = node
	list.Head.Next = node
}

func (list *List[K, V]) Remove(node *Node[K, V]) {
	if node == nil {
		return
	}

	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	node.Next = nil
	node.Prev = nil
}

func (list *List[K, V]) Back() *Node[K, V] {
	if list.Head.Next == list.Tail {
		return nil
	}

	return list.Tail.Prev
}

func (list *List[K, V]) MoveToFront(node *Node[K, V]) {
	if node == nil {
		return
	}

	list.Remove(node)
	list.PushFront(node)
}
