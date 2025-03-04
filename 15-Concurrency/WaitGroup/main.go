package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

var TodoLit = []string{}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 1; i <= 10; i++ {
		// wg.Add(1)
		GetTodo(i, &wg)
	}
	wg.Wait()
	fmt.Println(TodoLit)
}

func GetTodo(id int, wg *sync.WaitGroup) {
	GetURl("https://jsonplaceholder.typicode.com/todos/"+strconv.Itoa(id), wg)
}

func GetURl(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	responcebody, err := io.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		panic(err)
	}
	TodoLit = append(TodoLit, string(responcebody))
}
