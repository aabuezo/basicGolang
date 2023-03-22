package main

import "fmt"

func main() {
	c := make(chan int)

	// send
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("sending...")
			c <- i
		}
		close(c)
	}()

	// receive
	for v := range c {
		fmt.Println("received: ", v)
	}

	fmt.Println("finished")
}
