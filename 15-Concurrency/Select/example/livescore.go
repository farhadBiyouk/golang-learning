package example

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// https://footba11.co/json/liveFeed
// https://livescore-api.varzesh3.com/v1.0/livescore/today

// referer: https://www.varzesh3.com/

func SelectExample() {
	resource1 := make(chan string)
	resource2 := make(chan string)

	go GetResponse(resource2, "https://footba11.co/json/liveFeed")
	go GetResponse(resource1, "https://footba11.co/json/liveFeed")

	select {
	case result1 := <-resource1:
		println(result1)
	case result2 := <-resource2:
		println(result2)

	}
	PrintlnWithTime("End ")
}

func GetResponse(content chan<- string, url string) {
	client := http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	request.Header = http.Header{}

	request.Header.Add("referer", "https://www.varzesh3.com/")
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	destination := &bytes.Buffer{}
	if err = json.Indent(destination, responseBody, "", "  "); err != nil {
		panic(err)
	}
	PrintlnWithTime("Before set content")
	content <- destination.String()
	PrintlnWithTime("After set content")

}

func PrintlnWithTime(args ...any) {
	fmt.Printf("Time: %s, %v\n", time.Now().Format(time.RFC3339Nano), args)
}
