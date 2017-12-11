package main

import "testing"

// Prevents compiler optimizing out the loop.
var result int

const str1 = "hello"
const str2 = "world"

func MultiCopies(temp []byte) (len int) {
	temp = temp[:0]
	len += copy(temp[len:], str1[:3])
	len += copy(temp[len:], str2)
	return len
}

func Concat(temp []byte) (len int) {
	temp = temp[:0]
	len += copy(temp[len:], str1[:3]+str2)
	return len
}

func BenchmarkCopies(b *testing.B) {
	buf := make([]byte, len(str1)+len(str2))
	var temp int
	for i := 0; i < b.N; i++ {
		temp = MultiCopies(buf)
	}
	result = temp
}

func BenchmarkConcat(b *testing.B) {
	buf := make([]byte, len(str1)+len(str2))
	var temp int
	for i := 0; i < b.N; i++ {
		temp = Concat(buf)
	}
	result = temp
}
