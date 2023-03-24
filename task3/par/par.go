package par

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

const DataFile = "lorem.txt"

type wordFrequency struct {
	word  string
	count int
}

// Return the word frequencies of the text argument.
func WordCount(text string) map[string]int {
	freqs := make(map[string]int)
	words := strings.Fields(text)

	wg := new(sync.WaitGroup)
	results := make(chan wordFrequency, len(words))

	for _, word := range words {
		wg.Add(1)
		go func(word string) {
			defer wg.Done()
			cleanWord := strings.Trim(strings.ToLower(word), ".,!")
			results <- wordFrequency{word: cleanWord, count: 1}
		}(word)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for wf := range results {
		freqs[wf.word] += wf.count
	}

	return freqs
}

// Benchmark how long it takes to count word frequencies in text numRuns times.
//
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
	data, err := os.ReadFile("lorem.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", WordCount(string(data)))

	numRuns := 100
	runtimeMillis := benchmark(string(data), numRuns)
	printResults(runtimeMillis, numRuns)
}
