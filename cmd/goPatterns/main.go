package main

import (
	"fmt"
	"sync"

	"github.com/dziedzicgrzegorz/goPatterns/internal/design/creational"
	"github.com/dziedzicgrzegorz/goPatterns/internal/samplehash"
)

func main() {
	samplehash.MakeHash("Grzegorz Dziedzic")

	var wg sync.WaitGroup
	const numGoroutines = 10
	wg.Add(numGoroutines)

	results := make(chan *creational.Single, numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			results <- creational.GetInstance()
		}()
	}

	wg.Wait()
	close(results)
	for result := range results {
		fmt.Printf("Instance address: %p\n", result)
	}

}
