package main

import (
	"fmt"
	"math/rand"
	"time"
)

// func from 1-boring has an extra parameter
// that being a string channel
// passes data into the string channel and sleeps for a random duration in ms
func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	// create an instance of a string channel
	c := make(chan string)

	// placeholder routine
	go boring("boring!", c)

	// prints data received from a channel
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("Leaving")
}
