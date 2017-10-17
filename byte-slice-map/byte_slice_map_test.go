package main

import (
	"crypto/md5"
	crand "crypto/rand"
	"flag"
	"hash/crc32"
	"log"
	"math/rand"
	"testing"
)

var keySz = flag.Int("key-size", 32, "length of byte slice used for the key")
var nKeys = flag.Int("num-keys", 30, "number of keys per map")

func byteBuffer() []byte {
	return make([]byte, *keySz)
}

func StringMap(buffer []byte) (map[string]struct{}, []string) {
	out := make(map[string]struct{}, *nKeys)
	keys := make([]string, *nKeys)
	for i := range keys {
		if _, err := crand.Read(buffer); err != nil {
			log.Fatal(err)
		}
		keys[i] = string(buffer)
		out[keys[i]] = struct{}{}
	}
	return out, keys
}

func StringMapAccess(lookup map[string]struct{}, keys []string) struct{} {
	var out struct{}
	order := rand.Perm(len(keys))
	for _, idx := range order {
		out = lookup[keys[idx]]
	}
	return out
}

func MD5Map(buffer []byte) (map[[16]byte]struct{}, [][16]byte) {
	out := make(map[[16]byte]struct{}, *nKeys)
	keys := make([][16]byte, *nKeys)
	for i := range keys {
		if _, err := crand.Read(buffer); err != nil {
			log.Fatal(err)
		}
		keys[i] = md5.Sum(buffer)
		out[keys[i]] = struct{}{}
	}
	return out, keys
}

func MD5MapAccess(lookup map[[16]byte]struct{}, keys [][16]byte) struct{} {
	var out struct{}
	order := rand.Perm(len(keys))
	for _, idx := range order {
		out = lookup[keys[idx]]
	}
	return out
}

func CRC32Map(buffer []byte, tab *crc32.Table) (map[uint32]struct{}, []uint32) {
	out := make(map[uint32]struct{}, *nKeys)
	keys := make([]uint32, *nKeys)
	for i := range keys {
		if _, err := crand.Read(buffer); err != nil {
			log.Fatal(err)
		}
		keys[i] = crc32.Checksum(buffer, tab)
		out[keys[i]] = struct{}{}
	}
	return out, keys
}

func CRC32MapAccess(lookup map[uint32]struct{}, keys []uint32) struct{} {
	var out struct{}
	order := rand.Perm(len(keys))
	for _, idx := range order {
		out = lookup[keys[idx]]
	}
	return out
}

// Prevents compiler optimizing out the procedure calls.
var result struct{}

func BenchmarkStringMap(b *testing.B) {
	buffer := byteBuffer()
	var temp struct{}
	for i := 0; i < b.N; i++ {
		temp = StringMapAccess(StringMap(buffer))
	}
	result = temp
}

func BenchmarkMD5Map(b *testing.B) {
	buffer := byteBuffer()
	var temp struct{}
	for i := 0; i < b.N; i++ {
		temp = MD5MapAccess(MD5Map(buffer))
	}
	result = temp
}

func BenchmarkCRC32Map(b *testing.B) {
	buffer := byteBuffer()
	// Used for crc32 checksum.
	table := crc32.MakeTable(crc32.IEEE)
	var temp struct{}
	for i := 0; i < b.N; i++ {
		temp = CRC32MapAccess(CRC32Map(buffer, table))
	}
	result = temp
}
