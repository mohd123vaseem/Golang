package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {

	list := []string{"apple", "mango", "orange"} //slice

	fmt.Println("list : ", list)

	list = append(list, "kiwi") //add new ele
	fmt.Println("list : ", list)

	new_list := append(list[1:3]) //last index is excluded
	// new_list := append(list[1:])//from one to end
	// new_list := append(list[:3])// from start to <3

	fmt.Println("new_list : ", new_list)

	//IMPORTANT
	scores := make([]int, 3)

	scores[0] = 15
	scores[1] = 11
	scores[2] = 13
	// scores[3]=14//this gives err bcoz slice is of 3 size

	fmt.Println("scores : ", scores)

	//but append works
	scores = append(scores, 17, 10)
	fmt.Println("scores after append : ", scores)

	/*append() is a built-in function whose job is to expand the slice when it runs out of room:

	It checks if the underlying array has enough spare capacity (cap) for the new items (15 and 16).

	Since scores is completely full (len = 3, cap = 3), append() allocates a brand-new, larger underlying array in memory (typically doubling the capacity to 6).

	It copies [11, 12, 13] into the new array, places 15 and 16 right after them, and updates len to 5.

	Finally, it returns the updated slice header, which you save back into scores
	*/

	//SORTING
	sort.Ints(scores)
	//or
	slices.Sort(scores)
	slices.Reverse(scores)
	fmt.Println("sorted scores : ", scores)

	//HOW TO REMOVE ELEMENTS BASED ON INDEX.

	nums := []string{"vas", "kas", "VASEEM", "naz", "mohd"}

	fmt.Println("nums : ", nums)

	index := 2

	nums = append(nums[:index], nums[index+1:]...)

	fmt.Println("after removal nums : ", nums)

	/*
		append()'s 2nd parameter (onward): Expects individual elements (elem1, elem2, ...), not a slice.

		nums[index+1:]: Produces a sub-slice containing everything from index+1 to the end.

		... (Unpacking operator): Takes that sub-slice and unpacks its items into individual comma-separated 
		arguments so append() can add them one by one.
	*/
}
