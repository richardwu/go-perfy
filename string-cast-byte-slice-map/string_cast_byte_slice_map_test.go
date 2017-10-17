package main

import (
	crand "crypto/rand"
	"flag"
	"log"
	"math/rand"
	"testing"
)

var keySz = flag.Int("key-size", 32, "length of byte slice used for the key")
var nKeys = flag.Int("num-keys", 30, "number of keys per map")

func byteBuffer() []byte {
	return make([]byte, *keySz)
}

func genKeys() [][]byte {
	buffer := byteBuffer()
	keys := make([][]byte, *nKeys)
	for i := range keys {
		if _, err := crand.Read(buffer); err != nil {
			log.Fatal(err)
		}
		keys[i] = make([]byte, *keySz)
		copy(keys[i], buffer)
	}
	return keys
}

func DirectIndexMap(keys [][]byte) map[string]struct{} {
	out := make(map[string]struct{}, len(keys))
	for _, key := range keys {
		out[string(key)] = struct{}{}
	}
	return out
}

func DirectIndexAccess(lookup map[string]struct{}, keys [][]byte) struct{} {
	var out struct{}
	order := rand.Perm(len(keys))
	for _, idx := range order {
		out = lookup[string(keys[idx])]
	}
	return out
}

func CastFirstMap(keys [][]byte) map[string]struct{} {
	out := make(map[string]struct{}, len(keys))
	var strKey string
	for _, key := range keys {
		strKey = string(key)
		out[strKey] = struct{}{}
	}
	return out
}

func CastFirstAccess(lookup map[string]struct{}, keys [][]byte) struct{} {
	var out struct{}
	var strKey string
	order := rand.Perm(len(keys))
	for _, idx := range order {
		strKey = string(keys[idx])
		out = lookup[strKey]
	}
	return out
}

// Prevents compiler optimizing out the procedure calls.
var result struct{}

func BenchmarkDirectIndex(b *testing.B) {
	var temp struct{}
	keys := genKeys()
	for i := 0; i < b.N; i++ {
		lookup := DirectIndexMap(keys)
		temp = DirectIndexAccess(lookup, keys)
	}
	result = temp
}

func BenchmarkCastFirst(b *testing.B) {
	var temp struct{}
	keys := genKeys()
	for i := 0; i < b.N; i++ {
		lookup := CastFirstMap(keys)
		temp = CastFirstAccess(lookup, keys)
	}
	result = temp
}
