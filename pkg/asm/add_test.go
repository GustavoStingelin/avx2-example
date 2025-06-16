package asm

import (
	"testing"
	"testing/quick"
)

//go:generate go run add_amd64.go -out add_amd64.s -stubs add_stub.go

func TestAdd(t *testing.T) {
	expect := func(x, y uint64) uint64 { return x + y }
	if err := quick.CheckEqual(Add, expect, nil); err != nil {
		t.Fatal(err)
	}
}
