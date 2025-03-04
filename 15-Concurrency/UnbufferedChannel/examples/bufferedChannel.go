package examples

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type RequestResponse struct {
	Url          string
	ResponseBody string
	Index        int
}

func FetchDataFromChannelBuffered() {
	start := time.Now()

	wg := sync.WaitGroup{}
	content := make(chan RequestResponse, 10)
	// sender := make(chan <- int)
	// receiver := make(<-chan int)

	wg.Add(10)
	for i := 0; i < 10; i++ {
		content <- RequestResponse{Url: "https://jsonplaceholder.typicode.com/todos/" + strconv.Itoa(i+1), Index: i + 1}
		go GetHttpRequestBuffered(content, &wg)
	}

	wg.Wait()
	close(content)

	for j := 0; j < 10; j++ {
		result := <-content
		fmt.Println(result)
	}
	fmt.Println(time.Since(start).Milliseconds())
}

func GetHttpRequestBuffered(content chan RequestResponse, wg *sync.WaitGroup) {
	defer wg.Done()
	result := <-content
	response, err := http.Get(result.Url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content <- RequestResponse{Url: result.Url, ResponseBody: string(responseBody), Index: result.Index}
}
