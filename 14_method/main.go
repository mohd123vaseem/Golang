package main

import "fmt"

type User struct {
	Name  string // Exported (accessible outside package)
	Age   int
	Email string

	jobId int // 1st letter smallcase so ,its private ,cannot be used outside the package
}

func (u User) get_jobId() { //this is method as it belongs to User struct

	fmt.Println("jobID is :", u.jobId)
}

func (u *User) changeAge() { //this is method as it belongs to User struct
         
	u.Age++
	return
}

func main() {

	vaseem := User{"vaseem", 25, "vaseem@gmail.com", 232} // Allowed here because main() is in the same package as struct
	fmt.Println("job id is :",vaseem.jobId)
	fmt.Println("vaseem :",vaseem)

	vaseem.get_jobId()
	vaseem.changeAge()
	fmt.Println("new age is : ",vaseem.Age)
}

/*
A receiver between the func keyword and the function name makes a function a method on a struct.

In Go, structs don't have functions written inside { ... } like C++ classes or Java classes. Instead,
you define a regular function outside the struct and attach it using a receiver:
*/
