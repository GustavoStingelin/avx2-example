package gostd

import (
	"crypto/sha256"
	"testing"
	"testing/quick"
)

func TestSum256(t *testing.T) {
	goNative := sha256.Sum256([]byte("HODL"))
	asmNative := Sum256([]byte("HODL"))
	if goNative != asmNative {
		t.Errorf("Expected %x, got %x", goNative, asmNative)
	}
}

func TestSum256Quick(t *testing.T) {
	f := func(data []byte) bool {
		if len(data) == 0 {
			return true // skip empty data
		}
		goNative := sha256.Sum256(data)
		asmNative := Sum256(data)
		return goNative == asmNative
	}
	if err := quick.Check(f, nil); err != nil {
		t.Errorf("Quick check failed: %v", err)
	}
}
