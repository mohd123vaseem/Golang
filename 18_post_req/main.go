package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const url = "https://jsonplaceholder.typicode.com/posts"

func main() {
	reqBody := strings.NewReader(`{
		"title": "Go POST Request",
		"body": "Learning how to send JSON in Go",
		"userId": 1
	}`)

	response, err := http.Post(url, "application/json", reqBody)
	if err != nil {
		panic(err) // response might be nil; exit safely before touching response.Body
	}
	defer response.Body.Close() // Runs only when response is valid and non-nil

	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("StatusCode is: %d\n", response.StatusCode)
	fmt.Printf("Response is: %s\n", string(content))
}

/*
If http.Post fails and returns response == nil, putting defer response.Body.Close() before if err != nil accesses .
Body on a nil pointer right then and crashes your program with a nil pointer panic before main() even reaches the end.

What Actually Happens Step-by-Step
defer evaluates response.Body right away:
Go sees defer response.Body.Close(), evaluates response.
Body to find the receiver to pass to .Close(), and panics immediately when response is nil.

Even if response weren't nil on an error:
When a network request fails with a non-nil err, there is either no response body or nothing to read/close, and panic(err) exits the program anyway.
*/
