package main

import "fmt"

func main() {
	// 1. Declare with default zero values: [0 0 0]
	var numbers [3]int //it cant be resized anymore
	numbers[0] = 10
	numbers[2] = 20

	// 2. Declare and initialize on one line
	colors := [3]string{"Red", "Green", "Blue"}

	// 3. Let the compiler count the size for you using '...'
	scores := [...]int{95, 88, 72, 100} // Compiler sets type to [4]int

	fmt.Println("values of array : ",numbers)      // Output: [10 20 0]
	fmt.Println("values of array : ",colors)       // Output: [Red Green Blue]
	fmt.Println("size of array : ",len(scores))  // Output: 4
}

/* Why You Rarely Use Arrays Directly in Go
Because arrays have a rigid fixed size and copy all their data when passed around, 
Go developers rarely use raw arrays in everyday code. Instead, they use Slices 
([]int—without a number inside the brackets), which are dynamic, resizable views built 
on top of arrays (similar to std::vector or std::span in C++).*/