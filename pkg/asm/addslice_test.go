package asm

//go:generate go run addslice_amd64.go -out addslice_amd64.s -stubs addslice_stub.go

import (
	"fmt"
	"math"
	"testing"
	"testing/quick"
)

func TestAddSlice(t *testing.T) {
	f := func(a, b []uint64) bool {
		if len(a) != len(b) {
			return true
		}
		dst := make([]uint64, len(a))
		AddSlice(dst, a, b)
		for i := range a {
			if dst[i] != a[i]+b[i] {
				return false
			}
		}
		return true
	}
	if err := quick.Check(f, nil); err != nil {
		t.Fatal(err)
	}
}

func TestAddSlice_ManualCases(t *testing.T) {
	cases := []struct {
		a, b   []uint64
		expect []uint64
	}{
		{[]uint64{1, 2, 3}, []uint64{4, 5, 6}, []uint64{5, 7, 9}},
		{[]uint64{0, 0}, []uint64{0, 0}, []uint64{0, 0}},
		{[]uint64{math.MaxUint64}, []uint64{3}, []uint64{2}}, // overflow case if uint wraps
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			dst := make([]uint64, len(c.a))
			AddSlice(dst, c.a, c.b)
			for j := range dst {
				if dst[j] != c.expect[j] {
					t.Fatalf("expected %v, got %v", c.expect, dst)
				}
			}
		})
	}
}
