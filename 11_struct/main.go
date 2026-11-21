package main

import "fmt"

// User defines a custom data structure.
// Note: There is no INHERITANCE IN GOLANG.
type User struct {
	Name  string
	Age   int
	Email string
}

// Pointer receiver (*User): modifies the actual struct
func (u *User) HaveBirthday() {
	u.Age++
}

// Value receiver (User): modifies only a temporary copy
func (u User) TryFakeBirthday() {
	u.Age++
}

func main() {
	// Positional initialization
	vaseem := User{"vaseem", 25, "vaseem@gmail.com"}

	fmt.Println("vaseem details : ", vaseem)
	fmt.Println("vaseem details : ", vaseem.Name, vaseem.Age)
	fmt.Printf("vaseem details : %+v\n", vaseem)

	// Safe: order doesn't matter, and omitted fields default to zero values
	u1 := User{
		Name: "Vaseem",
		Age:  25,
		// Email is omitted -> automatically ""
	}

	// Risky if User struct changes later: relies strictly on field order
	u2 := User{"Ali", 30, "ali@example.com"}

	fmt.Printf("u1: %+v\n", u1)
	fmt.Printf("u2: %+v\n", u2)

	// Testing value vs pointer receiver on vaseem
	vaseem.TryFakeBirthday()
	fmt.Println("After value receiver:", vaseem.Age) // Still 25

	vaseem.HaveBirthday() // Go automatically passes &vaseem
	fmt.Println("After pointer receiver:", vaseem.Age) // Now 26!
}