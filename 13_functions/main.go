package main

import "fmt"

func main() {//main function is the entry point of the program

	fmt.Println("this is main function")
	
	sum,st := solve(3,5)
	fmt.Println("sum is : ",sum,",message : ",st)

	sum2 :=proSum(1,2,3)
	fmt.Println("sum2 : ",sum2)
	sum3 := proSum(1,2,3,4)
	fmt.Println("sum3 : ",sum3)
	


}

func solve(a int,b int) (int ,string) {//3rd int is return type
	
	fmt.Println("this is the solve function")

	return a+b , "this is solve method"
}

func proSum(values ...int) int {//we can recive any no of ele so we store them in slice "values"
	
    sum :=0

	for i :=0; i<len(values); i++ {

		sum += values[i]
	}

	return sum
}