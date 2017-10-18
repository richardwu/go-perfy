package main

import (
	"flag"
	"math/rand"
	"testing"
)

var setSz = flag.Int("set-size", 50, "number of elements inserted into the set")
var maxElem = flag.Int("max-elem", 100, "the maximum element in the set (min is 1)")

func genSetElems(scratch []int) []int {
	for i := range scratch {
		scratch[i] = rand.Intn(*maxElem + 1)
	}
	return scratch
}

func SliceSet(elems []int) []int {
	slice := make([]int, 0, len(elems))
	for _, elem := range elems {
		newElem := true
		for _, prev := range slice {
			if prev == elem {
				newElem = false
				break
			}
		}
		if newElem {
			slice = append(slice, elem)
		}
	}
	return slice
}

func SliceSetAccess(slice []int, query []int) bool {
	var found bool
	for _, q := range query {
		for _, elem := range slice {
			if elem == q {
				found = true
				break
			}
		}
	}
	return found
}

func HashSet(elems []int) map[int]struct{} {
	hset := make(map[int]struct{}, len(elems))
	for _, elem := range elems {
		hset[elem] = struct{}{}
	}
	return hset
}

func HashSetAccess(hset map[int]struct{}, query []int) bool {
	var found bool
	for _, q := range query {
		_, found = hset[q]
	}
	return found
}

// Prevents compiler optimizing out the procedure calls.
var result bool

func BenchmarkSliceSet(b *testing.B) {
	var temp bool
	scratch := make([]int, *setSz)
	for i := 0; i < b.N; i++ {
		set := genSetElems(scratch)
		slice := SliceSet(set)
		query := genSetElems(scratch)
		temp = SliceSetAccess(slice, query)
	}
	result = temp
}

func BenchmarkHashSet(b *testing.B) {
	var temp bool
	scratch := make([]int, *setSz)
	for i := 0; i < b.N; i++ {
		set := genSetElems(scratch)
		hset := HashSet(set)
		query := genSetElems(scratch)
		temp = HashSetAccess(hset, query)
	}
	result = temp
}
