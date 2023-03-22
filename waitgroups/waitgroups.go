// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

const MOD = 100000

func main() {

	// sin goroutines
	start := time.Now()
	fmt.Println("CPU:\t\t", runtime.NumCPU())
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())
	foo2()
	bar()
	elapsed := time.Since(start)
	fmt.Printf("took: %s\n\n", elapsed)

	// con goroutines
	start = time.Now()

	wg.Add(3)
	go foo(1)
	go foo(2)
	go foo(3)
	bar()

	fmt.Println("CPU:\t\t", runtime.NumCPU())
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())

	wg.Wait()

	elapsed = time.Since(start)
	fmt.Printf("took: %s\n", elapsed)
}

func foo(n int) {
	for i := 0; i < 1000000; i++ {
		if i%MOD == 0 {
			fmt.Printf("foo(%d): %d\n", n, i/MOD)
		}
	}
	wg.Done()
}

func foo2() {
	for i := 0; i < 1000000; i++ {
		if i%MOD == 0 {
			fmt.Printf("foo: %d\n", i/MOD)
		}
	}
}

func bar() {
	for i := 0; i < 1000000; i++ {
		if i%MOD == 0 {
			fmt.Printf("bar: %d\n", i/MOD)
		}
	}
}
