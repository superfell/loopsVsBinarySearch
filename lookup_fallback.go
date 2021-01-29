// +build !amd64

package loopVsBinarySearch

// Lookup returns the index into the array 'x' of the value 'k', or -1 if its not there. If k appears at multiple locations, you'll get one of them as the return value, it may not be the first one.
// This version is used in builds for targets other than amd64
func Lookup(k byte, x *[16]byte) int32 {
	for i := 0; i < 16; i++ {
		if x[i] == k {
			return int32(i)
		}
	}
	return int32(-1)
}
