// +build ignore

package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

func main() {
	TEXT("Lookup", NOSPLIT, "func(k byte, x *[16]byte) int32")
	Pragma("noescape")
	Doc("Lookup returns the index into the array 'x' of the value 'k', or -1 if its not there." +
		" If k appears at multiple locations, you'll get one of them as the return value, it may not be the first one.")
	x := Load(Param("x"), GP64())
	k := Load(Param("k"), GP32())

	xKey := XMM()
	MOVD(k, xKey)
	VPBROADCASTB(xKey, xKey) // xmm register now contains the value k in all 16 bytes

	xArr := XMM()
	MOVUPD(Mem{Base: x}, xArr) // xmm register now contains the 16 bytes of the array x

	// Compare bytes for equality between the 2 xmm registers.
	// xArr is updated with the result. Where they're equal the byte is set to FF
	// otherwise its set to 0
	PCMPEQB(xKey, xArr)

	rv := GP64()
	rOffset := GP64()
	XORQ(rOffset, rOffset)       // resOffset = 0
	MOVQ(xArr, rv)               // get the lower 8 bytes from the xmm register into rv
	TESTQ(rv, rv)                // is rv 0? if not, at least one byte was equal
	JNZ(LabelRef("returnCount")) // jump to converting that back to a index

	MOVHLPS(xArr, xArr) // move top 64 bits to lower 64 bits in xmm register
	MOVQ(xArr, rv)      // move lower 8 bytes into rv
	TESTQ(rv, rv)
	JZ(LabelRef("notFound")) // is rv 0? if so there's no matches, so return -1
	// the match was found in the top 8 bytes, so we need
	// to offset the final calculated index by 8.
	MOVQ(U64(8), rOffset)

	Label("returnCount") // return tailing zeros / 8 + offset
	idx := GP64()
	TZCNTQ(rv, idx)    // set idx to the number of trailing zeros in rv
	SHRQ(Imm(3), idx)  // divide idx by 8 to get from bit position to byte posn.
	ADDQ(rOffset, idx) // add the result offset in.

	Store(idx.As32(), ReturnIndex(0)) // return the final index as the result.
	RET()

	Label("notFound")
	rMiss := GP32()
	MOVL(U32(0xFFFFFFFF), rMiss)
	Store(rMiss, ReturnIndex(0)) // return -1
	RET()

	Generate()
}
