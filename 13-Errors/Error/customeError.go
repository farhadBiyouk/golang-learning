package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type HttpError struct {
	StatusCode int
	Message    string
}

func main() {
	res, err := GetRequest("")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)

}

func (error *HttpError) Error() string {
	return fmt.Sprintf("an error has occurred %d - %s", error.StatusCode, error.Message)

}

func NewHttpError(statusCode int, msg string) *HttpError {
	return &HttpError{StatusCode: statusCode, Message: msg}
}

func GetRequest(url string) (string, error) {
	if len(url) == 0 {
		return "", NewHttpError(400, "Bad request ")
	}
	response, err := http.Get(url)
	if err != nil {
		return "", NewHttpError(500, "bad server error")
	}
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", NewHttpError(200, "response have error")
	}
	return string(responseBody), nil
}
