package main

import (
	"fmt"
	"math"
	"time"
)

// N
const NUMBER = 10000000

// LEN set to 0 to start appending primes at the very beginning
const LEN = 0

// CAPACITY set to 700k because there are no more than ~7% of primes in N
const CAPACITY = NUMBER * .7

var listOfPrimes []int = make([]int, LEN, CAPACITY)

func main() {

	fmt.Println("Prime numbers in ", NUMBER)
	fmt.Println("Golang")
	start := time.Now()

	FillListOfPrimes(NUMBER)

	elapsed := time.Since(start)
	fmt.Printf("took: %s\n", elapsed)

	PrintListOfPrimes(20)
	fmt.Printf("Total primes in %d: %d\n", NUMBER, len(listOfPrimes))
}

// IsPrime checks if a specific number is prime or not
func IsPrime(number int) bool {

	// special case since 2 is the only even prime number
	if number == 2 {
		return true
	}

	sqr := int(math.Sqrt(float64(number)))
	for i := 2; i <= sqr; i++ {
		if (number % i) == 0 {
			return false
		}
	}
	return true
}

// FillListOfPrimes fills listOfPrimes with all primes in (number)
func FillListOfPrimes(number int) {
	for i := 1; i <= number; i++ {
		if IsPrime(i) {
			listOfPrimes = append(listOfPrimes, i)
		}
	}
}

// PrintListOfPrimes prints the first (n) primes in listOfPrimes
func PrintListOfPrimes(n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d, ", listOfPrimes[i])
	}
	fmt.Println("...")
}
