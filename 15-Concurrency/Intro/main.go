package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	value := 0 // shared resource
	go task1()
	go task2()
	go task3()

	go func() {
		value++
	}()
	go func() {
		value += 1
	}()
	go func() {
		value += 3
	}()
	fmt.Println(value)

	time.Sleep(time.Second)
}

func task1() {
	fmt.Println("task1")
}
func task2() {
	fmt.Println("task2")
}
func task3() {
	fmt.Println("task3")
}
