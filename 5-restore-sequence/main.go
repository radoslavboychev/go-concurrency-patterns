package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	str  string
	wait chan bool
}

// fanIn accepts any number of channels of type Message as an argument
// returns a resulting channel of type Message
// containing all the data from the argument channels
func fanIn(inputs ...<-chan Message) <-chan Message {
	c := make(chan Message)
	for i := range inputs {
		input := inputs[i]
		go func() {
			for {
				c <- <-input
			}
		}()
	}
	return c
}

// boring accepts a string to pass into a channel
// creates an object of type Message and passes it into the resulting channel
func boring(msg string) <-chan Message {
	c := make(chan Message)
	waitChan := make(chan bool)

	// goroutine that creates a new object of type Message containing the message and a wait channel
	// and feeds it into the resulting channel
	go func() {
		for i := 0; ; i++ {
			c <- Message{
				str:  fmt.Sprintf("%s %d", msg, i),
				wait: waitChan,
			}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			// wait for the value to be received
			<-waitChan
		}
	}()

	// return a channel to caller
	return c
}

func main() {

	// create a new channel and feed 2 channels into it
	c := fanIn(boring("Joe"), boring("Ahn"))

	//
	for i := 0; i < 5; i++ {
		msg1 := <-c
		fmt.Println(msg1.str)
		msg2 := <-c
		fmt.Println(msg2.str)

		// each goroutine has to wait
		msg1.wait <- true
		msg2.wait <- true
	}

	fmt.Println("Leaving...")
}
