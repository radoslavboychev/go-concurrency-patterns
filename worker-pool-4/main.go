package main

import (
	"fmt"
	"log"
	"net/http"
)

type Site struct {
	URL string
}

type Result struct {
	URL    string
	Status int
}


func crawl(wId int, jobs <-chan Site, results chan<- Result) {
	for site := range jobs {
		log.Printf("Worker ID: %d\n", wId)
		resp, err := http.Get(site.URL)
		if err != nil {
			log.Println(err.Error())
		}
		results <- Result{Status: resp.StatusCode, URL: resp.Request.URL.String()}
	}
}


func main() {
	fmt.Println("worker pools in Go")

	jobs := make(chan Site, 3)
	result := make(chan Result, 3)

	for w := 1; w <= 3; w++ {
		go crawl(w, jobs, result)
	}
	

	urls := []string{
		"https://google.com",
		"https://youtube.com",
		"https://example.com",
		"https://twitter.com",
	}

	for _, url := range urls {
		jobs <- Site{URL: url}
	}
	close(jobs)

	for a := 1; a <= 4; a++ {
		result := <-result
		log.Println(result)
	}

}
