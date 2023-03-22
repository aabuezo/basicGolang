package main

import (
	"fmt"
	"math"
	"time"
)

const LEN = 0

// there are no more than ~7% of primes in N
const CAPACITY = 700000

var listOfPrimes []int = make([]int, LEN, CAPACITY)

func main() {

	// N
	number := 10000000
	fmt.Println("Prime numbers in ", number)
	fmt.Println("Golang")
	start := time.Now()

	FillListOfPrimes(number)

	elapsed := time.Since(start)
	fmt.Printf("took: %s\n", elapsed)

	PrintListOfPrimes(20)
	fmt.Printf("Total primes in %d: %d\n", number, len(listOfPrimes))
}

// IsPrime checks if a specific number is prime or not
func IsPrime(number int) bool {

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
