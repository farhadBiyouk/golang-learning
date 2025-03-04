package main

import (
	"fmt"
	"sync"
)

func SimpleGood() {
	counter := 0 // shared resource  thread safe -> race condition -> deed lock acquire
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	wg.Add(1000000)
	for i := 0; i < 1_000_000; i++ {
		go func() {
			defer wg.Done()
			defer mu.Unlock()

			mu.Lock()
			counter += 1
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}

func SimpleBad() {
	counter := 0 // shared resource  thread safe -> race condition -> deed lock acquire
	wg := sync.WaitGroup{}
	wg.Add(1000000)
	for i := 0; i < 1_000_000; i++ {
		go func() {
			defer wg.Done()
			counter += 1
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
