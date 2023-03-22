package main

import (
	"testing"
)

func TestIsPrime(t *testing.T) {
	primes := []int{1, 2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	not_primes := []int{4, 6, 8, 9, 10, 12, 14, 15, 16, 18, 20, 21, 22, 24, 25, 26, 27, 28, 30}

	for _, v := range primes {
		if !IsPrime(v) {
			t.Errorf("Expected true but got false with: %d\n", v)
		}
	}

	for _, v := range not_primes {
		if IsPrime(v) {
			t.Errorf("Expected false but got true with: %d\n", v)
		}
	}
}
