package loopVsBinarySearch

import (
	"math/rand"
	"testing"
)

var rnd = rand.New(rand.NewSource(42))
var keys = make([]byte, 16)

func init() {
	for i := 0; i < len(keys); i++ {
		keys[i] = byte(i)
	}
	rnd.Shuffle(len(keys), func(i, j int) {
		keys[i], keys[j] = keys[j], keys[i]
	})
}

func Benchmark_Loop(b *testing.B) {
	n := new(nodeLoop)
	for _, k := range keys {
		n.put(k, int(k))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, k := range keys {
			v := n.get(k)
			if v != int(k) {
				b.Errorf("Got unexpected value of %d for key %d", v, k)
			}
		}
	}
}
func Benchmark_LoopOneBoundsCheck(b *testing.B) {
	n := new(nodeLoop)
	for _, k := range keys {
		n.put(k, int(k))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, k := range keys {
			v := n.get2(k)
			if v != int(k) {
				b.Errorf("Got unexpected value of %d for key %d", v, k)
			}
		}
	}
}

func Benchmark_LoopRev(b *testing.B) {
	n := new(nodeLoop)
	for _, k := range keys {
		n.put(k, int(k))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, k := range keys {
			v := n.getReverse(k)
			if v != int(k) {
				b.Errorf("Got unexpected value of %d for key %d", v, k)
			}
		}
	}
}
func Benchmark_LoopRevOneBoundsCheck(b *testing.B) {
	n := new(nodeLoop)
	for _, k := range keys {
		n.put(k, int(k))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, k := range keys {
			v := n.getReverse2(k)
			if v != int(k) {
				b.Errorf("Got unexpected value of %d for key %d", v, k)
			}
		}
	}
}

func Benchmark_BinarySearch(b *testing.B) {
	n := new(nodeSorted)
	for _, k := range keys {
		n.put(k, int(k))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, k := range keys {
			v := n.get(k)
			if v != int(k) {
				b.Errorf("Got unexpected value of %d for key %d", v, k)
			}
		}
	}
}

func Benchmark_BinarySearchOneBoundsCheck(b *testing.B) {
	n := new(nodeSorted)
	for _, k := range keys {
		n.put(k, int(k))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, k := range keys {
			v := n.get2(k)
			if v != int(k) {
				b.Errorf("Got unexpected value of %d for key %d", v, k)
			}
		}
	}
}

func Benchmark_BinarySearchInlined(b *testing.B) {
	n := new(nodeSorted)
	for _, k := range keys {
		n.put(k, int(k))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, k := range keys {
			v := n.getInlinedBinSearch(k)
			if v != int(k) {
				b.Errorf("Got unexpected value of %d for key %d", v, k)
			}
		}
	}
}

func Benchmark_BinarySearchInlinedOneBoundsCheck(b *testing.B) {
	n := new(nodeSorted)
	for _, k := range keys {
		n.put(k, int(k))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, k := range keys {
			v := n.getInlinedBinSearch2(k)
			if v != int(k) {
				b.Errorf("Got unexpected value of %d for key %d", v, k)
			}
		}
	}
}
