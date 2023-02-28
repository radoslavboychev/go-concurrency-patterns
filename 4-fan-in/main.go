package main

import (
	"fmt"
	"math/rand"
	"time"
)

// boring takes in a message string, fills it in with data and returns it
func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

// fanIn accepts 2 channels as arguments, reads data from them using goroutines
// into a results channel and returns it
func fanIn(c1, c2 <-chan string) <-chan string {
	c := make(chan string)

	// reads value from channel c1 and sends it to the results channel
	go func() {
		for {
			v1 := <-c1
			c <- v1
		}

	}()

	// reads value from channel c2 and sends it to the results channel
	go func() {
		for {
			c <- <-c2
		}
	}()

	// return the resulting channel
	return c
}

// fanInSimple accepts any amount of channels as arguments
// feeds all of their data into a resulting channel
// returns the resulting channel
func fanInSimple(cs ...<-chan string) <-chan string {
	c := make(chan string)
	for _, ci := range cs {
		go func(cv <-chan string) {
			for {
				c <- <-cv
			}
		}(ci)
	}
	return c
}

func main() {
	// c := fanIn(boring("Joe"),boring("Ahn"))
	c := fanInSimple(boring("Joe"), boring("Ahn"))

	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("Leaving")
}
