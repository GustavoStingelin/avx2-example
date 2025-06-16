package main

import (
	"avx2-example/pkg/asm"
	"avx2-example/pkg/gostd"
	"testing"
)

func BenchmarkAddSlice(b *testing.B) {
	const n = 1024
	a := make([]uint64, n)
	bSlice := make([]uint64, n)
	dst := make([]uint64, n)
	for i := 0; i < n; i++ {
		a[i] = uint64(i)
		bSlice[i] = uint64(i)
	}
	b.ResetTimer()
	b.Run("asm", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			asm.AddSlice(dst, a, bSlice)
		}
	})
	b.Run("gostd", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			gostd.AddSlice(dst, a, bSlice)
		}
	})
}
