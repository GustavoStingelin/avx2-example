//go:build ignore

package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

func main() {
	// Define the function signature for assembly generation
	TEXT("AddSlice", NOSPLIT, "func(dst, a, b []uint64)")
	Doc("AddSlice adds corresponding elements of a and b into dst.")

	// Load base addresses of slices into registers
	dst := Load(Param("dst").Base(), GP64()) // destination slice
	a := Load(Param("a").Base(), GP64())     // slice a
	b := Load(Param("b").Base(), GP64())     // slice b
	n := Load(Param("a").Len(), GP64())      // number of elements (len(a))

	// Calculate how many full 4-element (4*64-bit = 32 bytes) AVX2 operations we can do
	vecN := GP64()
	MOVQ(n, vecN)
	SHRQ(U8(2), vecN) // vecN = n / 4

	// Compute remainder to handle tail elements
	rem := GP64()
	MOVQ(n, rem)
	ANDQ(U8(3), rem) // rem = n % 4

	// If no full vector operations are needed, jump to tail processing
	CMPQ(vecN, Imm(0))
	JE(LabelRef("tail"))

	// Allocate a YMM register for AVX2 processing
	vec := YMM()

	// Vectorized addition loop: add 4 uint64s at a time
	Label("loop")
	VMOVDQU(Mem{Base: a}, vec)     // Load 4 uint64s from a into vec
	VPADDQ(Mem{Base: b}, vec, vec) // Add 4 uint64s from b to vec
	VMOVDQU(vec, Mem{Base: dst})   // Store the result into dst
	ADDQ(U8(32), a)                // Move a pointer forward by 32 bytes (4 * 8)
	ADDQ(U8(32), b)                // Move b pointer forward by 32 bytes
	ADDQ(U8(32), dst)              // Move dst pointer forward by 32 bytes
	DECQ(vecN)                     // Decrement loop counter
	JNZ(LabelRef("loop"))          // Loop if more iterations remain

	// Handle remaining (non-vectorizable) elements
	Label("tail")
	MOVQ(rem, vecN)
	CMPQ(vecN, Imm(0))
	JE(LabelRef("done")) // If no remainder, skip to done

	tmp := GP64() // Temporary register to hold computation

	// Scalar loop: add remaining 1-3 elements
	Label("loop_tail")
	MOVQ(Mem{Base: a}, tmp)    // Load uint64 from a
	ADDQ(Mem{Base: b}, tmp)    // Add corresponding uint64 from b
	MOVQ(tmp, Mem{Base: dst})  // Store result in dst
	ADDQ(U8(8), a)             // Move a pointer forward by 8 bytes
	ADDQ(U8(8), b)             // Move b pointer forward by 8 bytes
	ADDQ(U8(8), dst)           // Move dst pointer forward by 8 bytes
	DECQ(vecN)                 // Decrement counter
	JNZ(LabelRef("loop_tail")) // Continue until all tail elements are processed

	// Cleanup
	Label("done")
	VZEROUPPER() // Clear upper parts of YMM registers to avoid AVX-SSE transition penalties
	RET()        // Return from function

	Generate() // Finalize code generation
}
