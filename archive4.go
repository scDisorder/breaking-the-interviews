package main

import "fmt"

func bufferChannelCheck() {
	intCh := make(chan int, 3)
	intCh <- 10 // intCh cap=3, len=1
	intCh <- 3  // intCh cap=3, len=2
	intCh <- 24 // intCh cap=3, len=3
	// locked state - func main waits until there is a free space in a channel
	// adding (uncommenting) fmt.Println(<-intCh) at L12 before will fix that deadlock
	//fmt.Println(<-intCh)
	//intCh <- 15

	fmt.Println(<-intCh)
	fmt.Println("The End")
}
