package main

import (
	"fmt"
	"math"
	"time"
)

const LEN = 0
const CAPACITY = 100000

var listOfPrimes []int = make([]int, LEN, CAPACITY)

func main() {

	number := 10000000
	fmt.Println("Prime numbers in ", number)
	fmt.Println("Golang")
	start := time.Now()

	FillListOfPrimes(number)

	elapsed := time.Since(start)
	fmt.Printf("took: %s\n", elapsed)

	PrintListOfPrimes(20)
	fmt.Printf("Total primes in %d: %d\n", number, GetListOfPrimesSize())
}

func IsPrime(number int) bool {

	num := int(math.Sqrt(float64(number)))
	for i := 2; i <= num; i++ {
		if (number % i) == 0 {
			return false
		}
	}
	return true
}

func FillListOfPrimes(number int) {
	for i := 1; i <= number; i++ {
		if IsPrime(i) {
			listOfPrimes = append(listOfPrimes, i)
		}
	}
}

func PrintListOfPrimes(num int) {
	for i := 0; i < num; i++ {
		fmt.Printf("%d, ", listOfPrimes[i])
	}
	fmt.Println("...")
}

func GetListOfPrimesSize() int {
	return len(listOfPrimes)
}
