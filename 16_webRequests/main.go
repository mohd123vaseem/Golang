package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const url = "https://jsonplaceholder.typicode.com/todos/1"

func main() {

	fmt.Println("Web Requests")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Response data type is: %T \n", response)

	defer response.Body.Close() //in go close connections yourself. As they dont close automatically.

	databytes, err := io.ReadAll(response.Body)//it reutrn data in byte form

	var resp strings.Builder

	resp.Write(databytes)//it writes the databytes into resp

	if err!=nil {

		panic(err)
	}

	fmt.Println(" #Data is :", string(databytes))
	fmt.Println(" ##Data is :", resp.String())
}
