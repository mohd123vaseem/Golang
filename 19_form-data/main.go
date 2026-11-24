package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const formURL = "https://httpbin.org/post"

func main() {
	data := url.Values{}
	data.Add("firstname", "hitesh")
	data.Add("lastname", "choudhary")
	data.Add("email", "hitesh@go.dev")

	response, err := http.PostForm(formURL, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("status code is:", response.StatusCode)
	fmt.Println("Response is:", string(content))
}

/*
url is the net/url package, and url.
Values is a map (map[string][]string) in Go.
*/