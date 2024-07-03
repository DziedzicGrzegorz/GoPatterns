package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/dziedzicgrzegorz/goPatterns/internal/design/creational"
	"github.com/dziedzicgrzegorz/goPatterns/internal/samplehash"
	"go.uber.org/zap"
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
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	sugar := logger.Sugar()
	defer logger.Sync()

	sugar.Debug("this is a debug message")
	sugar.Info("this is an info message")
	sugar.Warn("this is a warn message")
	sugar.Error("this is an error message")
	sugar.Fatal("this is a fatal message")

}
