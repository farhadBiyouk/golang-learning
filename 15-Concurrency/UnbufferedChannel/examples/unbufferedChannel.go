package examples

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

func FetchDataFromChannel() {
	start := time.Now()

	content := make(chan string)

	for i := 0; i < 10; i++ {
		go GetHttpRequest(content, "https://jsonplaceholder.typicode.com/todos/", i)
	}

	for j := 0; j < 10; j++ {
		result := <-content
		fmt.Println(result)
	}
	fmt.Println(time.Since(start).Milliseconds())
}

func GetHttpRequest(content chan<- string, BaseUrl string, index int) {
	response, err := http.Get(BaseUrl + "/" + strconv.Itoa(index))

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content <- string(responseBody)
}
