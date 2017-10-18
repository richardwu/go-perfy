package main

import (
	"bytes"
	"crypto/rand"
	"flag"
	"log"
	"testing"
)

var numElems = flag.Int("num-elems", 100, "number of elements in input")
var sz = flag.Int("elem-size", 32, "length of each []byte element")

func genElem(scratch []byte) []byte {
	if _, err := rand.Read(scratch); err != nil {
		log.Fatal(err)
	}
	return scratch
}

func UnallocSlice(scratch []byte) []byte {
	var out []byte
	for i := 0; i < *numElems; i++ {
		out = append(out, genElem(scratch)...)
	}
	return out
}

func PreallocSlice(scratch []byte) []byte {
	out := make([]byte, 0, *numElems*(*sz))
	for i := 0; i < *numElems; i++ {
		out = append(out, genElem(scratch)...)
	}
	return out
}

func BytesBuffer(scratch []byte) []byte {
	var out bytes.Buffer
	for i := 0; i < *numElems; i++ {
		if _, err := out.Write(genElem(scratch)); err != nil {
			log.Fatal(err)
		}
	}
	return out.Bytes()
}

// Prevents compiler optimizing out the procedure calls.
var result []byte

func BenchmarkUnallocSlice(b *testing.B) {
	scratch := make([]byte, *sz)
	var temp []byte
	for i := 0; i < b.N; i++ {
		temp = UnallocSlice(scratch)
	}
	result = temp
}

func BenchmarkPreallocSlice(b *testing.B) {
	scratch := make([]byte, *sz)
	var temp []byte
	for i := 0; i < b.N; i++ {
		temp = PreallocSlice(scratch)
	}
	result = temp
}

func BenchmarkBytesBuffer(b *testing.B) {
	scratch := make([]byte, *sz)
	var temp []byte
	for i := 0; i < b.N; i++ {
		temp = BytesBuffer(scratch)
	}
	result = temp
}
