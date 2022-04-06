package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"sync"
	"time"
)

const DataFile = "loremipsum.txt"

// Return the word frequencies of the text argument.
func WordCount(text string) map[string]int {
	freqs := make(map[string]int)
	text = strings.ToLower(text)
	words := strings.Fields(text)
	ch := make(chan map[string]int)
	var wg sync.WaitGroup
	length := len(words)
	size := 3000
	for i, j := 0, size; i < length; i, j = j, j+size {

		if j > length {
			j = length
		}
		wg.Add(1)
		go func(words []string) {
			local_freqs := make(map[string]int)
			for _, word := range words {
				word = strings.Trim(word, ".")
				word = strings.Trim(word, ",")
				local_freqs[word]++
			}
			ch <- local_freqs
			wg.Done()
		}(words[i:j])
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	for sub := range ch {
		for word, val := range sub {
			freqs[word] = val
		}
	}
	return freqs
}

// Benchmark how long it takes to count word frequencies in text numRuns times.
// Return the total time elapsed.
func benchmark(text string, numRuns int) int64 {
	start := time.Now()
	for i := 0; i < numRuns; i++ {
		WordCount(text)
	}
	runtimeMillis := time.Since(start).Nanoseconds() / 1e6

	return runtimeMillis
}

// Print the results of a benchmark
func printResults(runtimeMillis int64, numRuns int) {
	fmt.Printf("amount of runs: %d\n", numRuns)
	fmt.Printf("total time: %d ms\n", runtimeMillis)
	average := float64(runtimeMillis) / float64(numRuns)
	fmt.Printf("average time/run: %.2f ms\n", average)
}

func main() {
	// read in DataFile as a string called data
	data, err := ioutil.ReadFile("/Users/David/Documents/INDA_folder/dpeilitz-palinda-3/singleworker/loremipsum.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	//fmt.Printf("%#v: \n", WordCount(string(data), wg))
	numRuns := 100
	runtimeMillis := benchmark(string(data), numRuns)
	printResults(runtimeMillis, numRuns)
}
