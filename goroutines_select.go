package main

import "fmt"

func fibonacci(c, quit chan int) {
	fmt.Println("Sender: Starting channel sender")
	x, y := 0, 1
	for {
		fmt.Println("Sender: Put ", x)
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Sender: quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		fmt.Println("Receiver: Starting channel receiver...")
		for i := 0; i < 10; i++ {
			fmt.Println("Receiver: Got ", <-c)
		}
		fmt.Println("Receiver: Sending end signal...")
		quit <- 0
	}()
	fibonacci(c, quit)
}

