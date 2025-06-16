package main

import (
	"avx2-example/pkg/gostd"
	"bytes"
	"crypto/sha256"
	"testing"
)

var benchData = bytes.Repeat([]byte("HODL"), 256) // 1024 bytes of data

func BenchmarkSHA256(b *testing.B) {
	b.Run("asm", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sha256.Sum256(benchData)
		}
	})
	b.Run("stdgo", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			gostd.Sum256(benchData)
		}
	})
}
