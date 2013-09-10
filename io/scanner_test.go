package io

import (
	"fmt"
	"testing"
)

func TestScan(t *testing.T) {
	fmt.Printf("Testing scan...\n")
	const in, out = "/Users/mark/Projects/testing/flickr", 4
	if x := Scan(in); len(x) != out {
		t.Errorf("scan(%v) = %v, want %v", in, x, out)
	}
}
