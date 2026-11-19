package main

import (
	"fmt"
)

func main() {

	list := []string{"apple", "mango", "orange"} //slice

	fmt.Println("list : ", list)

	list = append(list, "kiwi") //add new ele
	fmt.Println("list : ", list)

	new_list := append(list[1:3])//last index is excluded
    // new_list := append(list[1:])//from one to end
    // new_list := append(list[:3])// from start to <3

	fmt.Println("new_list : ",new_list)

}
