package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result string
type Search func(query string) Result

var (
	Web   = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result %q\n", kind, query))
	}
}

func Google(query string) []Result {
	c := make(chan Result)

	go func() {
		c <- Web(query)
	}()

	go func() {
		c <- Image(query)
	}()

	go func() {
		c <- Video(query)
	}()

	var results []Result

	// when a timeout is passed
	timeout := time.After(500 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case r := <-c:
			results = append(results, r)
		case <-timeout:
			fmt.Println("timeout")
			return results
		}
	}
	return results

}

func main() {
	rand.NewSource(time.Now().UnixNano())
	start := time.Now()
	results := Google("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}
