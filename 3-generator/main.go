package main

import (
	"fmt"
	"math/rand"
	"time"
)

// boring returns a string channel containing the string passed as argument
func boring(msg string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
		// close the channel
		close(c)
	}()

	return c //return the channel to the caller

}

func main() {
	joe := boring("Joe")
	ahn := boring("Ahn")

	// print the data from the channels
	for i := 0; i < 10; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-ahn)
	}

	fmt.Println("Leaving...")

}
