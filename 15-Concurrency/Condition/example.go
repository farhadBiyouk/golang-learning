package main

import (
	"fmt"
	"sync"
	"time"
)

var UserList = []int{}
var ready = false

func main() {
	condition := sync.NewCond(&sync.Mutex{})
	for i := 0; i < 60; i++ {
		go NewRequest(i, condition)
	}
	time.Sleep(time.Second * 5)
}

func NewRequest(userId int, condition *sync.Cond) {
	fmt.Println("New request")
	Checking(userId, condition)
	defer condition.L.Unlock()
	condition.L.Lock()
	for !ready {
		condition.Wait()
	}
	println("User ", userId, "start streaming")
}
func Checking(userId int, condition *sync.Cond) {
	fmt.Println(userId, "waiting for start streaming")
	defer condition.L.Unlock()
	condition.L.Lock()
	UserList = append(UserList, userId)
	if len(UserList) == 55 {
		ready = true
		condition.Broadcast()
		fmt.Println("End waiting lets go to start streaming ")
	}

}
