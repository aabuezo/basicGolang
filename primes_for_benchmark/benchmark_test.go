package main

import "testing"

func BenchmarkFillListOfPrimes(b *testing.B) {
	number := 10000000
	for i := 0; i < b.N; i++ {
		FillListOfPrimes(number)
	}
}
