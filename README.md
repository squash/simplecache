package simplecache // import "github.com/squash/simplecache"

Package simplecache is a naive implementation or a LRU cache, with a soft
maximum intended to keep the cache between 90 and 100% full. it is designed
to be easy to use and understand, and fast enough for most uses.

TYPES

type LRU struct {
	// Has unexported fields.
}
    LRU is an instance of our LRU model cache

func NewLRU(size, age int) *LRU
    NewLRU creates a new cache object with a soft maximum size.

func (l *LRU) Dump() []interface{}
    Return the contents of the cache

func (l *LRU) Fetch(key string) (interface{}, bool)
    Fetch grabs an entry from the cache

func (l *LRU) Flush()
    Flush cache entirely

func (l *LRU) Peek(key string) (interface{}, bool)
    Peek grabs the entry without updating it's last access time.

func (l *LRU) Store(key string, value interface{})
    Store adds or replaces an entry in the cache

