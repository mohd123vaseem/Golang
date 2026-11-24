package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println(myurl)

	// Parsing the URL
	result, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}

	fmt.Println("Scheme:", result.Scheme)
	fmt.Println("Host:", result.Host)
	fmt.Println("Path:", result.Path)
	fmt.Println("Port:", result.Port())
	fmt.Println("RawQuery:", result.RawQuery)

	// Extracting query parameters into url.Values (key-value map)
	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)
	fmt.Println("Course name:", qparams["coursename"])
	fmt.Println("Payment ID:", qparams["paymentid"])

	for _, val := range qparams {
		fmt.Println("Param is:", val)
	}

	// Constructing a new URL from parts (always pass a reference &url.URL{} , to avoid copying thr struct and)
	/*The other main reason is type compatibility—many Go functions and methods expect a pointer (*url.URL),
	  not a plain struct value (url.URL).*/
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		// RawPath: "user=hitesh",
		RawQuery: "user=hitesh",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
