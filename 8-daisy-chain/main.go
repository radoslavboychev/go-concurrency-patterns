package main

import "fmt"

// accepts 2 channels as arguments, takes the value from the right,
// adds 1 to it, passes it into the left
func f(left, right chan int) {
	left <- 1 + <-right
}

// 
func main() {
	const n = 1000
	leftmost := make(chan int)
	left := leftmost
	right := leftmost

	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}

	go func(c chan int) {
		c <- 1
	}(right)
	fmt.Println(<-leftmost)
}
