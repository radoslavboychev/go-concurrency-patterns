package main

import (
	"fmt"
	"sync"
	"time"
)

type JobInput struct {
	startTime time.Time
	endTime   time.Time
}

// WORKER POOLS PATTERN
// 1. Job producer feeds the input to a buffered channel
// 2. A number of workers are created
// 3. Each worker completes a business functionality job and outputs the results to a channel
// 4. 

// represents the individual response from each business functionality job
// job responses are combined in the end in a singular Response type
type JobResponse struct{}

// represents the output of a business functionality job, whatever that type is
type Response struct {
	finalOutput string
}

// individual job to be ran that performs a business functionality
func businessFunctionalityJob(jobInput JobInput) JobResponse {
	fmt.Println("Executing job...")
	return JobResponse{}
}

// receives a list of responses from individual jobs and combines them into a single one
func combineResponses(jobResponses []JobResponse) Response {
	return Response{finalOutput: "Well done!!"}
}

// for every channel with job input, feeds the response into an output channel
func startWorker(jobInputChan <-chan JobInput, jobOutputChan chan<- JobResponse, wg *sync.WaitGroup) {
	defer wg.Done()
	for jobInput := range jobInputChan {
		jobOutputChan <- businessFunctionalityJob(jobInput)
	}
}

// accept a receiving channel
func businessFunctionality(jobInputChan chan<- JobInput) {
	jobInputs := []JobInput{JobInput{}, JobInput{}}
	for _, jobInput := range jobInputs {
		jobInputChan <- jobInput
	}
	close(jobInputChan)
}

func main() {
	// set the number of job workers
	num_workers := 3

	// set the buffer number for job channels
	// (amount of messages they can receive before blocking)
	num_buffer := 10

	// create a buffered channel
	jobsChan := make(chan JobInput, num_buffer)

	// set the channel for
	resultsChan := make(chan JobResponse, num_buffer)

	// WaitGroup to use for workers
	wg := sync.WaitGroup{}

	// start workers working jobs
	for i := 0; i < num_workers; i++ {
		go startWorker(jobsChan, resultsChan, &wg)
	}

	//
	go businessFunctionality(jobsChan)

	// create a variable to store all responses
	var responses []JobResponse

	// sync group to use when combining responses
	wgResp := sync.WaitGroup{}
	wgResp.Add(1)

	go func() {
		defer wgResp.Done()
		for resp := range resultsChan {
			responses = append(responses, resp)
		}
	}()

	wg.Wait()
	close(resultsChan)
	wgResp.Wait()

	fmt.Println(combineResponses(responses).finalOutput)
}
