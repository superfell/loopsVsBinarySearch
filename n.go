package loopVsBinarySearch

import (
	"sort"
)

type nodeLoop struct {
	key   [16]byte
	val   [16]interface{}
	count int
}

func (n *nodeLoop) put(k byte, v interface{}) {
	idx := n.count
	n.key[idx] = k
	n.val[idx] = v
	n.count++
}

func (n *nodeLoop) get(k byte) interface{} {
	for i := 0; i < n.count; i++ {
		if n.key[i] == k {
			return n.val[i]
		}
	}
	return nil
}

func (n *nodeLoop) get2(k byte) interface{} {
	_ = n.key[n.count-1]
	for i := 0; i < n.count; i++ {
		if n.key[i] == k {
			return n.val[i]
		}
	}
	return nil
}

func (n *nodeLoop) getReverse(k byte) interface{} {
	for i := n.count - 1; i >= 0; i-- {
		if n.key[i] == k {
			return n.val[i]
		}
	}
	return nil
}

func (n *nodeLoop) getReverse2(k byte) interface{} {
	_ = n.key[n.count-1]
	for i := n.count - 1; i >= 0; i-- {
		if n.key[i] == k {
			return n.val[i]
		}
	}
	return nil
}

type nodeSorted struct {
	key   [16]byte
	val   [16]interface{}
	count int
}

func (n *nodeSorted) put(k byte, v interface{}) {
	idx := sort.Search(n.count, func(i int) bool {
		return n.key[i] >= k
	})
	copy(n.key[idx+1:], n.key[idx:int(n.count)])
	copy(n.val[idx+1:], n.val[idx:int(n.count)])
	n.key[idx] = k
	n.val[idx] = v
	n.count++
}

func (n *nodeSorted) get(k byte) interface{} {
	idx := sort.Search(n.count, func(i int) bool {
		return n.key[i] >= k
	})
	if idx < int(n.count) && n.key[idx] == k {
		return n.val[idx]
	}
	return nil
}

func (n *nodeSorted) get2(k byte) interface{} {
	_ = n.key[n.count-1]
	idx := sort.Search(n.count, func(i int) bool {
		return n.key[i] >= k
	})
	if idx < int(n.count) && n.key[idx] == k {
		return n.val[idx]
	}
	return nil
}

func (n *nodeSorted) getInlinedBinSearch(k byte) interface{} {
	// impl of sort.Search manually inlined here
	i, j := 0, int(n.count)
	for i < j {
		h := int(uint(i+j) >> 1) // avoid overflow when computing h
		// i ≤ h < j
		if n.key[h] < k {
			i = h + 1 // preserves f(i-1) == false
		} else {
			j = h // preserves f(j) == true
		}
	}
	// i == j, f(i-1) == false, and f(j) (= f(i)) == true  =>  answer is i.
	if i < int(n.count) && n.key[i] == k {
		return n.val[i]
	}
	return nil
}

func (n *nodeSorted) getInlinedBinSearch2(k byte) interface{} {
	// impl of sort.Search manually inlined here
	_ = n.key[n.count-1]
	i, j := 0, int(n.count)
	for i < j {
		h := int(uint(i+j) >> 1) // avoid overflow when computing h
		// i ≤ h < j
		if n.key[h] < k {
			i = h + 1 // preserves f(i-1) == false
		} else {
			j = h // preserves f(j) == true
		}
	}
	// i == j, f(i-1) == false, and f(j) (= f(i)) == true  =>  answer is i.
	if i < int(n.count) && n.key[i] == k {
		return n.val[i]
	}
	return nil
}
