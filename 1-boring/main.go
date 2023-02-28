package main

import (
	"fmt"
	"math/rand"
	"time"
)

// function that prints a string and the loop index
func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

// most basic usage of goroutines
func main() {

	// call the goroutine
	go boring("boring!") //spawns a goroutine

	// lets the main goroutine run forever
	// for {
	// }

	// ends the main goroutine 2 seconds after starting
	fmt.Println("Listening")
	time.Sleep(2 * time.Second)
	fmt.Println("Leaving")

}
