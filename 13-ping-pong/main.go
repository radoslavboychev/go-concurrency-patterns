package main

import (
	"fmt"
	"time"
)

// Ball represents a ball in a ping pong match
type Ball struct {
	hits int
}

// A game has a name and
func game(playerName string, table chan *Ball) {
	for {
		ball := <-table
		ball.hits++
		fmt.Println(playerName, ball.hits)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}

func main() {

	table := make(chan *Ball)

	go game("ping", table)
	go game("pong", table)

	table <- new(Ball)

	time.Sleep(2 * time.Second)
	<-table
	panic("show stack")

}
