package loopVsBinarySearch

import "testing"

func Benchmark_UnrolledLoop(b *testing.B) {
	n := new(nodeLoop)
	for _, k := range keys {
		n.put(k, int(k))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, k := range keys {
			v := n.unrolledLoop(k)
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
		if n.get(99) != -1 {
			t.Errorf("key 99 shouldn't have a value")
		}
	}
}

func Test_Masks2(t *testing.T) {
	n := new(nodeMasks)
	for i := 0; i < 8; i++ {
		n.put(byte(i), i)
	}
	n.put(8, 8)
	for i := 0; i < 9; i++ {
		v := n.get(byte(i))
		if v != i {
			t.Errorf("Got unexpected result of %d for key %d", v, i)
		}
	}
}
