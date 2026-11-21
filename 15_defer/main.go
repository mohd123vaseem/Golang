package main

import "fmt"

func main() {

	defer fmt.Println("world")
	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("Hello")

	myDefer()
}

func myDefer()  {
	
	for i:=0; i<5; i++ {

		defer fmt.Println(i)
	}
}

// defer schedules the call to run right before the enclosing function returns (in LIFO order)
//in case of mutiple defer they follow LIFO like a stack.