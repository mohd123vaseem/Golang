package main

import "fmt"

func main() {

	mp := make(map[int]string)

	mp[1] = "abc"
	mp[2] = "xyz"
	mp[3] = "123"

	fmt.Println("map : ", mp)
	fmt.Println("map : ", mp[3])

	delete(mp, 3)
	fmt.Println("map : ", mp)

	mp[4] = "vaseem"

	for key, val := range mp {
		fmt.Printf("Key: %d, Value: %s\n", key, val)
	}

	for _, val := range mp { //igonring key
		fmt.Printf("Key: ignored, Value: %s\n", val)
	}

	/*
		In Go, a for ... range loop over a map iterates in a completely random, non-deterministic order.
	*/

	/*
	%d and %s are called format specifiers (or "verbs" in Go). They act as placeholders inside fmt.
	Printf() that tell Go what data type to expect and how to format it when printing.

	## Most Used `fmt` Verbs in Go

	* **`%v`** — Any value (default format for ints, strings, slices, maps, structs)
	* **`%+v`** — Struct with field names (`{Name:Alex Age:24}`)
	* **`%#v`** — Go-syntax representation (for debugging)
	* **`%T`** — Data type of the variable (`int`, `[]string`, `main.User`)
	* **`%s`** — String (`string` or `[]byte`)
	* **`%d`** — Base-10 integer (`int`, `int64`, `uint`)
	* **`%f`** — Floating-point number (`%.2f` for 2 decimal places)
	* **`%c`** — Single character (`rune` or `byte`)
	* **`%t`** — Boolean (`true` or `false`)
	* **`%p`** — Pointer / memory address in hex
	*/

}
