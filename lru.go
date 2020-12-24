// Package simplecache is a naive implementation or a LRU cache, with a soft maximum intended to keep the cache between 90 and 100% full.
// It is desined to be easy to use and understand, and fast enough for most uses.
package simplecache

import (
	"sort"
	"sync"
	"time"
)

// LRU is an instance of our LRU model cache
type LRU struct {
	entries   map[interface{}]cacheEntry
	lock      sync.Mutex
	sizelimit int
	ages      []age
	purgeto   int
}

type cacheEntry struct {
	value   interface{}
	lasthit time.Time
}

// NewLRU creates a new cache object with a soft maximum size.
func NewLRU(size int) *LRU {
	var l LRU
	l.sizelimit = size
	l.purgeto = int(float64(size) * 0.9)
	l.entries = make(map[interface{}]cacheEntry)

	return &l
}

// Store adds or replaces an entry in the cache
func (l *LRU) Store(key, value interface{}) {
	var e cacheEntry
	e.value = value
	e.lasthit = time.Now()
	l.lock.Lock()
	l.entries[key] = e
	l.collect()
	l.lock.Unlock()
}

// Fetch grabs an entry from the cache
func (l *LRU) Fetch(key interface{}) (interface{}, bool) {
	l.lock.Lock()
	e, ok := l.entries[key]
	if ok {
		e.lasthit = time.Now()
		l.entries[key] = e
	}
	l.lock.Unlock()
	return e.value, ok
}

// Peek grabs the entry without updating it's last access time.
func (l *LRU) Peek(key interface{}) (interface{}, bool) {
	l.lock.Lock()
	e, ok := l.entries[key]
	l.lock.Unlock()
	return e.value, ok
}

// Return the contents of the cache
func (l *LRU) Dump() []interface{} {
	var i []interface{}
	l.lock.Lock()
	for x := range l.entries {
		i = append(i, l.entries[x].value)
	}
	l.lock.Unlock()
	return i
}

// Delete a single entry by key
func (l *LRU) Delete(key interface{}) bool {
	l.lock.Lock()
	_, ok := l.entries[key]
	if ok {
		delete(l.entries, key)
	}
	l.lock.Unlock()
	return ok
}

// Count the number of entries in the cache
func (l *LRU) Count() int {
	l.lock.Lock()
	tmp := len(l.entries)
	l.lock.Unlock()
	return tmp
}

// Flush cache entirely
func (l *LRU) Flush() {
	l.lock.Lock()
	l.entries = make(map[interface{}]cacheEntry)
	l.lock.Unlock()
}

// ages is a second index sorted by age used to purge
type age struct {
	key     interface{}
	lasthit time.Time
}

// collect enforces size maximums, in chunks. We'll purge to 90% full when we run so we're not running too often.
func (l *LRU) collect() {
	if len(l.entries) < l.sizelimit {
		// We aren't full so let's not run
		return
	}
	var ages []age
	for x := range l.entries {
		var a age
		a.key = x
		a.lasthit = l.entries[x].lasthit
		ages = append(ages, a)
	}
	// Sort by age
	sort.Slice(ages, func(a, b int) bool { return ages[a].lasthit.After(ages[b].lasthit) })
	// Now we remove all entries from the 90% age mark up
	for x := l.purgeto; x < len(ages); x++ {
		delete(l.entries, ages[x].key)
	}
}
