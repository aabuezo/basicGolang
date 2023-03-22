// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// RACE CONDITION
	count := 0

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("GR:", runtime.NumGoroutine())
	gs := 100
	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := count
			// time.Sleep/time.Second)
			runtime.Gosched()
			v++
			count = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("GR:", runtime.NumGoroutine())
	fmt.Println("count:", count)
}
