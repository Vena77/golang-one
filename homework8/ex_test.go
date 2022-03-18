package main

import (
	"testing"
)

func init() {
	direct1 := "C:/Users/Acer 1/go/src/NewNew1"
	ListByReadDir(direct1)
}

func BenchmarkTest1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		thesimilar()
	}
}

func BenchmarkTest2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		compare1()
	}
}

func Example() {
	direct1 := "C:/Users/Acer 1/go/src/NewNew1"
	ListByReadDir(direct1)
}
