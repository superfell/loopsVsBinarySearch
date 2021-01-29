package loopVsBinarySearch

import (
	"testing"
)

func Benchmark_IndexByte(b *testing.B) {
	n := new(nodeLoop)
	for _, k := range keys {
		n.put(k, int(k))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, k := range keys {
			v := n.getIndexByte(k)
			if v != int(k) {
				b.Errorf("Got unexpected value of %d for key %d", v, k)
			}
		}
	}
}

func Benchmark_UnrolledLoop(b *testing.B) {
	n := new(nodeLoop)
	for _, k := range keys {
		n.put(k, int(k))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, k := range keys {
			v := n.getUnrolledLoop(k)
			if v != int(k) {
				b.Errorf("Got unexpected value of %d for key %d", v, k)
			}
		}
	}
}

func Benchmark_GetLookupAsm(b *testing.B) {
	n := new(nodeLoop)
	for _, k := range keys {
		n.put(k, int(k))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, k := range keys {
			v := n.getLookupAsm(k)
			if v != int(k) {
				b.Errorf("Got unexpected value of %d for key %d", v, k)
			}
		}
	}
}

func Benchmark_Masks(b *testing.B) {
	n := new(nodeMasks)
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
func Benchmark_MasksWithFinalLoop(b *testing.B) {
	n := new(nodeMasks)
	for _, k := range keys {
		n.put(k, int(k))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, k := range keys {
			v := n.getWithFinalLoop(k)
			if v != int(k) {
				b.Errorf("Got unexpected value of %d for key %d", v, k)
			}
		}
	}
}

func Benchmark_MasksWithBitTwiddling(b *testing.B) {
	n := new(nodeMasks)
	for _, k := range keys {
		n.put(k, int(k))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, k := range keys {
			v := n.getMoreBitTwiddling(k)
			if v != int(k) {
				b.Errorf("Got unexpected value of %d for key %d", v, k)
			}
		}
	}
}

func Test_Masks(t *testing.T) {
	n := new(nodeMasks)
	for i := 0; i < 16; i++ {
		k := byte(i) + 3
		v := 100 + i*2
		n.put(k, v)
		for j := 0; j <= i; j++ {
			k := byte(j) + 3
			v := 100 + j*2
			if n.get(k) != v {
				t.Errorf("Got unexpected result of %d for key %d after inserting %d keys", n.get(k), k, i+1)
			}
		}
		if n.get(99) != nil {
			t.Errorf("key 99 shouldn't have a value")
		}
	}
}

func Test_MasksBitTwiddling(t *testing.T) {
	n := new(nodeMasks)
	for i := 0; i < 16; i++ {
		n.put(byte(i), i)
	}
	for i := 0; i < 16; i++ {
		v := n.getMoreBitTwiddling(byte(i))
		if v != i {
			t.Errorf("Got unexpected result of %d for key %d", v, i)
		}
		v = n.getMoreBitTwiddling(byte(i) + 100)
		if v != nil {
			t.Errorf("Got unexpected result of %v for key %d", v, i+100)
		}
	}
}

func Test_Lookup(t *testing.T) {
	keys := [16]byte{11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26}
	for i := 0; i < 32; i++ {
		idx := Lookup(byte(i), &keys)
		if i >= 11 && i <= 26 {
			if int(idx) != i-11 {
				t.Errorf("lookup %d returned %d\n", i, idx)
			}
		} else if idx != -1 {
			t.Errorf("Lookup %d : return index %d\n", i, idx)
		}
	}
}
