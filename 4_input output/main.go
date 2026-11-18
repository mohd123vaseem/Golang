package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin);

	fmt.Println(" how was the Pizza")

	intput ,_ := reader.ReadString('\n')//here we ignored error using underscore(_)

	fmt.Println("thanks of the rating : ",intput) 

	var name string
	var age int

	fmt.Println("Enter you name and age please : ")
   
	// underscore(_) char is used to ignore returned value or error
	_ , err := fmt.Scanln(&name,&age)

	if err !=nil {

		fmt.Println("Invalid input")
	}

	fmt.Println("Thanks for the info Mr/Ms ",name ,"of age ",age)

}