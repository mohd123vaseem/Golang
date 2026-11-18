package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
    
	fmt.Println("Please rate this pizza")

	input,_ := reader.ReadString('\n')

	// num,err := strconv.ParseFloat(input,64)//Error
	num,err := strconv.ParseFloat(strings.TrimSpace(input),64)//Correct line

	if err!=nil{
        
		fmt.Println(err)
	}

	fmt.Println("Thanks for the rating : ",num+1);
    
	//TO convert string to integer
	var s string = "123"
 
	n,err := strconv.Atoi(s)
    
	fmt.Println("value is : ",n+2)

}