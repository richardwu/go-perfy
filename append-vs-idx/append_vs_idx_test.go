package main

import (
	"flag"
	"testing"
)

var sliceSz = flag.Int("slice-size", 32, "length of final slice")

func AppendSlice() []int {
	foo := make([]int, 0, *sliceSz)
	count := 0
	for count < *sliceSz {
		foo = append(foo, 1)
		count++
	}
	return foo
}

func IdxSlice() []int {
	foo := make([]int, *sliceSz)
	for i := range foo {
		foo[i] = 1
	}
	return foo
}

// Prevents compiler optimizing out the procedure calls.
var result []int

func BenchmarkAppend(b *testing.B) {
	var temp []int
	for i := 0; i < b.N; i++ {
		temp = AppendSlice()
	}
	result = temp
}

func BenchmarkIdx(b *testing.B) {
	var temp []int
	for i := 0; i < b.N; i++ {
		temp = IdxSlice()
	}
	result = temp
}
