package main

import "fmt"

func doubleValue(n *int) {
	*n = *n * 2 // Modifies the original variable directly
}

func main() {

	fmt.Println("lets start pointers")

	//Pointers

	var ptr *int      // nil pointer
	fmt.Println(*ptr) // PANIC: runtime error: invalid memory address or nil pointer dereference

	num := 10

	ptr2 := &num

	fmt.Println("address of ptr2 : ", ptr2)
	fmt.Println("value of ptr2 : ", *ptr2)

	//pass by refernce
	doubleValue(&num) // Pass the memory address of 'num'
	fmt.Println(num)  // Output: 10
}

/*Why Do We Use Pointers?
By default, Go is pass-by-value. When you pass a normal variable into a function,
Go makes a complete copy of it.

Without a pointer: Changing the variable inside a function only changes the copy,
leaving the original untouched.

With a pointer: Passing the memory address lets the function directly modify the original variable—and
avoids copying huge structs in memory, making your program faster. */
