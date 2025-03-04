package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Employee struct {
	ID     int
	Salary int64
}

func main() {

	wg := sync.WaitGroup{}
	// mu := sync.Mutex{}
	var AmountAccount int64 = 25_500_000_000
	employees := []Employee{}

	for i := 0; i < 5_000; i++ {
		employees = append(employees, Employee{ID: i, Salary: GetRandomNumber()})
	}

	wg.Add(len(employees))
	for _, item := range employees {
		go func() {
			defer wg.Done()

			// method1
			// defer mu.Unlock()
			// mu.Lock()
			// if item.Salary < AmountAccount {
			// 	AmountAccount -= item.Salary
			// }
			atomic.AddInt64(&AmountAccount, -item.Salary)
		}()
	}
	wg.Wait()
	fmt.Println(AmountAccount)
}

func GetRandomNumber() int64 {
	return 5_000_000
}
