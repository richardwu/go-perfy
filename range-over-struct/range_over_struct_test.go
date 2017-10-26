package main

import (
	"flag"
	"testing"
)

var numStructs = flag.Int("num-structs", 200, "number of structs to iterate over")

const bytes = 1048576

type largeStruct struct {
	data [bytes]byte
}

func createStructs() []largeStruct {
	out := make([]largeStruct, *numStructs)
	for i := range out {
		for j := range out[i].data {
			out[i].data[j] = byte(j)
		}
	}
	return out
}

func IndexRange(structs []largeStruct) [bytes]byte {
	var temp [bytes]byte
	for i := range structs {
		temp = structs[i].data
	}
	return temp
}

func CopyRange(structs []largeStruct) [bytes]byte {
	var temp [bytes]byte
	for _, struc := range structs {
		temp = struc.data
	}
	return temp
}

func IndexRangeEveryByte(structs []largeStruct) byte {
	var temp byte
	for i := range structs {
		for j := range structs[i].data {
			temp = structs[i].data[j]
		}
	}
	return temp
}

func CopyRangeEveryByte(structs []largeStruct) byte {
	var temp byte
	for _, struc := range structs {
		for _, byt := range struc.data {
			temp = byt
		}
	}
	return temp
}

// Prevents compiler optimizing out the procedure calls.
var result [bytes]byte
var resultEvery byte

func BenchmarkIndexRange(b *testing.B) {
	var temp [bytes]byte
	structs := createStructs()
	for i := 0; i < b.N; i++ {
		temp = IndexRange(structs)
	}
	result = temp
}

func BenchmarkCopyRange(b *testing.B) {
	var temp [bytes]byte
	structs := createStructs()
	for i := 0; i < b.N; i++ {
		temp = CopyRange(structs)
	}
	result = temp
}
func BenchmarkIndexRangeEveryByte(b *testing.B) {
	var temp byte
	structs := createStructs()
	for i := 0; i < b.N; i++ {
		temp = IndexRangeEveryByte(structs)
	}
	resultEvery = temp
}

func BenchmarkCopyRangeEveryByte(b *testing.B) {
	var temp byte
	structs := createStructs()
	for i := 0; i < b.N; i++ {
		temp = CopyRangeEveryByte(structs)
	}
	resultEvery = temp
}
