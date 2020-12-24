package simplecache

import (
	"fmt"
	"testing"
)

type testdata struct {
	word   string
	number int
}

var result interface{}

func TestCache(t *testing.T) {
	l := NewLRU(10)
	for x := 0; x <= 20; x++ {
		var d testdata
		d.number = x
		d.word = fmt.Sprintf("%d", x)
		l.Store(d.word, d)
	}
	// We should now have entries 11-20
	_, ok := l.Fetch("9") // should not exist
	if ok {
		t.Fatal("Cache entry 9 should have been purged")
	}
	_, ok = l.Fetch("12")
	if !ok {
		fmt.Println("Cache contents: %#V\n", l.Dump())
		t.Fatal("Cache entry 12 should exist")
	}

	var d testdata
	i, ok := l.Fetch("20")
	if !ok {
		t.Fatal("Cache entry 20 should exist")
	}
	d = i.(testdata)
	if d.number != 20 {
		t.Fatal("Cache entry 20 has wrong value")
	}
}

func BenchmarkAdds(b *testing.B) {
	l := NewLRU(b.N / 2)
	for n := 0; n < b.N; n++ {
		l.Store(fmt.Sprintf("%d", n), n)
	}
}

func BenchmarkFetches(b *testing.B) {
	var tmp interface{}
	l := NewLRU(b.N)
	for n := 0; n < b.N; n++ {
		l.Store(fmt.Sprintf("%d", n), n)
	}
	for n := 0; n < b.N; n++ {
		tmp, _ = l.Fetch(fmt.Sprintf("%d", n))
	}
	result = tmp
}
