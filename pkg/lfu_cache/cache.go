package lfucache

const defaultFreq = 1

type LFU[K comparable, V any] struct {
	cap         int
	minFreq     int
	keys        map[int]*Set[K]
	values      map[K]V
	frequencies map[K]int
}

func NewLFU[K comparable, V any](cap int) *LFU[K, V] {
	return &LFU[K, V]{
		cap:         cap,
		minFreq:     0,
		keys:        make(map[int]*Set[K]),
		values:      make(map[K]V),
		frequencies: make(map[K]int),
	}
}

func (lfu *LFU[K, V]) Get(key K) (val V, found bool) {
	if _, ok := lfu.values[key]; !ok {
		return val, false
	}

	val = lfu.values[key]
	freq := lfu.frequencies[key]
	lfu.updateState(key, val, freq)
	return val, true
}

func (lfu *LFU[K, V]) Put(key K, value V) {
	if _, ok := lfu.values[key]; ok {
		freq := lfu.frequencies[key]
		lfu.updateState(key, value, freq)
		return
	}

	if len(lfu.values) == lfu.cap {
		lfu.popWeak()
	}

	lfu.values[key] = value
	lfu.frequencies[key] = defaultFreq
	if _, ok := lfu.keys[defaultFreq]; !ok {
		lfu.keys[defaultFreq] = NewSet[K]()
	}
	lfu.keys[defaultFreq].Add(key)
	lfu.minFreq = defaultFreq
}

func (lfu *LFU[K, V]) updateState(key K, value V, freq int) {
	set := lfu.keys[freq]
	set.Remove(key)
	if set.IsEmpty() {
		delete(set.keys, key)
		if lfu.minFreq == freq {
			lfu.minFreq++
		}
	}

	lfu.values[key] = value
	newFreq := freq + 1
	lfu.frequencies[key] = newFreq
	if _, ok := lfu.keys[newFreq]; !ok {
		lfu.keys[newFreq] = NewSet[K]()
	}
	lfu.keys[newFreq].Add(key)
}

func (lfu *LFU[K, V]) popWeak() {
	set := lfu.keys[lfu.minFreq]
	key := set.PopAny()
	if set.IsEmpty() {
		delete(set.keys, key)
	}

	delete(lfu.values, key)
	delete(lfu.frequencies, key)
}
