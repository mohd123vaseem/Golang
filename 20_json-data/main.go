package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON video")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline.in", "hit123", nil},
	}

	// Compact JSON using json.Marshal
	finalJson, err := json.Marshal(lcoCourses)
	//First return ([]byte): The JSON-encoded bytes (convert with string(finalJson) or
	//%s in fmt.Printf to read it as text).
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)

	// Pretty-printed JSON using json.MarshalIndent
	indentedJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", indentedJson)
}

/*
json.Marshal converts your Go data (like a struct, slice, or map) into a JSON byte slice ([]byte)
so it can be sent over a network, written to a file, or printed as JSON.
*/

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"website": "LearnCodeOnline.in",
		"tags": ["web-dev","js"]
	}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)// Unmarshals (decodes/reads) JSON bytes Into the lcoCourse struct fields
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	// Some cases where you just want to add data to a key-value map
	var myOnlineData map[string]any// we used any bcoz we dont know the data-type of values
	// var myOnlineData map[string]interface// we used interface bcoz we dont know the data-type of values
	
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is: %T\n", k, v, v)
	}
}