package main

import "fmt"

// boring runs the function until data is sent into the
// quit channel, when it quits
func boring(msg string, quit chan string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s %d", msg, i):
			case <-quit:
				fmt.Println("clean up")
				quit <- "See you!"
				return
			}
		}
	}()
	return c
}

// the boring function runs until anything is sent to the quit channel
// after which it finishes
func main() {
	quit := make(chan string)
	c := boring("joe", quit)
	for i := 3; i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- "Bye"
	fmt.Println("Joe say:", <-quit)
}
