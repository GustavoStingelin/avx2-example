package main

import (
	"avx2-example/pkg/asm"
	"avx2-example/pkg/gostd"
	"testing"
)

func BenchmarkAdd(b *testing.B) {
	b.Run("asm", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			asm.Add(uint64(i), uint64(i))
		}
	})
	b.Run("gostd", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = gostd.Add(uint64(i), uint64(i))
		}
	})
}
