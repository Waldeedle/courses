package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

var mutex = &sync.Mutex{}

// comparing a sequential read of a csv with a concurrent read of a csv
func main() {
	start := time.Now()

	//sequential read
	sequentialRead()

	//concurrent read
	// concurrentRead()

	fmt.Printf("Time taken: %s\n", time.Since(start))
}

func sequentialRead() {
	// Open the csv file
	file, err := os.Open("1000k.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a buffered reader
	r := bufio.NewReader(file)

	// Read 1 line at a time
	records := []string{}
	for {
		record, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		records = append(records, record)
	}
}

func concurrentRead() {
	// Open the csv file
	file, err := os.Open("1000k.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a buffered reader
	r := bufio.NewReader(file)

	// Read the entire file into memory
	lines := make([]string, 0)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		lines = append(lines, line)
	}

	// Use a wait group to wait for all the workers to finish
	var wg sync.WaitGroup

	// Number of workers
	numWorkers := 10000

	// Create the workers with the wait group to process the lines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(i int) {
			// Each worker processes 1/numWorkers of the lines
			for j := i * len(lines) / numWorkers; j < (i+1)*len(lines)/numWorkers; j++ {
				// Just read the line, don't do anything with it
			}
			wg.Done()
		}(i)
	}

	// Wait for all the workers to finish
	wg.Wait()
}
