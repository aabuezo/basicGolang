// Primes fills a list (slice) of primes from 1 to N and calculates the time it takes
package main

import (
	"fmt"
	"math"
	"time"
)

// NUMBER find ell primes from 1 to NUMBER
const NUMBER = 10000000

// LEN set to 0 to start appending primes at the very beginning
const LEN = 0

// CAPACITY is set to (NUMBER * .7) because for N > 10M, the number of primes
// tend to be less than 7%. For smaller Ns there could be a little overhead
// due to the increment in capacity, though negligible for this test
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

	// no need to calculate sqr each time within the for loop
	sqr := int(math.Sqrt(float64(number)))
	for i := 2; i <= sqr; i++ {
		if (number % i) == 0 {
			return false
		}
	}
	return true
}

// FillListOfPrimes fills listOfPrimes with all primes from 1 to NUMBER
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
