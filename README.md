# Introduction

simplecache implements a LRU (evicts the least recently used entry) cache, which can be an efficient way to store data that would otherwise have to be found externally each time. A good example is storing objects loaded from a database, in particular when your object may require multiple queries and additional processing to assemble. While a database query may take tens or hundreds of millisenconds per query, simplecache can retrieve your populated object in a few hundred nanoseconds.

# Implementation

simplecache is a naive, thread safe implementation. It uses Go's map system to store data, and can use any type of data for both key and value. There are faster, more complex implementations of a LRU cache out there, but this was designed to be read as much as used.

# How to use

Since we use Go's interface type for both keys and values, you can use any sot of data with simplecache. However, when retrieveing an object, you will be responsible for casting it back to it's original type. See this simple example: https://play.golang.org/p/IZqWtXdain2

simplecache will maintain a maximum entry size and cull entries to 90% capacity when that maximum is reached by removing the last recently used objects.

Of course, it is up to you to keep your cache from going stale. When your application mutates an object, simply Store() it again and it will overwrite the previous version.

# GODOC

package simplecache // import "github.com/squash/simplecache"

Package simplecache is a naive implementation or a LRU cache, with a soft
maximum intended to keep the cache between 90 and 100% full. It is desined
to be easy to use and understand, and fast enough for most uses. If you initialize
with a max size of 0 the cache will not purge entries automatically.

TYPES

type LRU struct {
// Has unexported fields.
}
LRU is an instance of our LRU model cache

func NewLRU(size int) \*LRU
NewLRU creates a new cache object with a soft maximum size.

func (l \*LRU) Delete(key interface{}) bool
Delete a single entry by key

func (l \*LRU) Dump() []interface{}
Return the contents of the cache

func (l \*LRU) Fetch(key interface{}) (interface{}, bool)
Fetch grabs an entry from the cache

func (l \*LRU) Flush()
Flush cache entirely

func (l \*LRU) Peek(key interface{}) (interface{}, bool)
Peek grabs the entry without updating it's last access time.

func (l \*LRU) Store(key, value interface{})
Store adds or replaces an entry in the cache
