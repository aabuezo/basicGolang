package main

import "fmt"

func main() {
	c := make(chan int)

	// send
	go foo(c)

	// receive
	bar(c)

	fmt.Println("finished")
}

func foo(c chan<- int) {
	for i := 0; i < 5; i++ {
		fmt.Println("sending...")
		c <- 42
	}
}

func bar(c <-chan int) {
	for i := 0; i < 5; i++ {
		fmt.Println("receiving...")
		<-c
		fmt.Println("...received!")
	}
}
