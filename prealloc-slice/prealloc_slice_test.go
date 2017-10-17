package main

import (
	"flag"
	"math/rand"
	"testing"
)

var maxElems = flag.Int("max-elems", 32, "maximum number of elements in the final slice")
var appendPct = flag.Float64("append-*appendPct", 0.9, "percentage of maximum number of elements to incude in the final slice")

func randBoolSlice() []bool {
	out := make([]bool, *maxElems)
	for i := range out {
		out[i] = rand.Float64() < *appendPct
	}
	return out
}

func Append(input []bool) []int {
	var out []int
	for i, in := range input {
		if in {
			out = append(out, i)
		}
	}
	return out
}

func MaxAlloc(input []bool) []int {
	out := make([]int, 0, len(input))
	for i, in := range input {
		if in {
			out = append(out, i)
		}
	}
	return out
}

func ExactAlloc(input []bool) []int {
	count := 0
	for _, in := range input {
		if in {
			count++
		}
	}
	out := make([]int, count)
	idx := 0
	for i, in := range input {
		if in {
			out[idx] = i
			idx++
		}
	}
	return out
}

// Prevents compiler optimizing out the procedure calls.
var result []int

func BenchmarkAppend(b *testing.B) {
	var temp []int
	for i := 0; i < b.N; i++ {
		in := randBoolSlice()
		temp = Append(in)
	}
	tempult = temp
}

func BenchmarkExactAlloc(b *testing.B) {
	var temp []int
	for i := 0; i < b.N; i++ {
		in := randBoolSlice()
		temp = ExactAlloc(in)
	}
	tempult = temp
}

func BenchmarkMaxAlloc(b *testing.B) {
	var temp []int
	for i := 0; i < b.N; i++ {
		in := randBoolSlice()
		temp = MaxAlloc(in)
	}
	tempult = temp
}
