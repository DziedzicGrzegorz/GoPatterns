package creational

import (
	"reflect"
	"sync"
	"testing"
)

func TestGetInstanceByOneThread(t *testing.T) {
	want := GetInstance()
	for i := 0; i < 10; i++ {
		if got := GetInstance(); !reflect.DeepEqual(got, want) {
			t.Errorf("GetInstance() = %v, want %v", got, want)
		}
	}
}

func TestGetInstanceByGoroutines(t *testing.T) {
	want := GetInstance()
	var wg sync.WaitGroup
	const numGoroutines = 10
	wg.Add(numGoroutines)

	results := make(chan *Single, numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			results <- GetInstance()
		}()
	}

	wg.Wait()
	close(results)

	for got := range results {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("GetInstance() = %v, want %v", got, want)
		}
	}
}
