package main

import "fmt"

const Id = 123456789//public is made by 1st letter as uppercase . IMPORTANT
//If 1st letter is smallcase then its private

func main() {

	fmt.Println("diff types of Data Types")
    
	//EXPLICIT WAY 

	//String
	var name string="vaseem"
	var s string // empty string
	fmt.Println("string s : ",s)
	fmt.Println("My name is ",name)
	fmt.Printf("my name is :%T \n" , name);
    
	//Bool
	var flag bool = true
	fmt.Println("falg value : ",flag)

	//Integer
	var num int = -123
	var num2 uint = 222//no sign allowed
	fmt.Println("num is : ",num,",",num2)
	var n int//it initialises the var with 0
	fmt.Println("n is zero bcoz golang dont put garbage val : ",n)

	//Float
	var fnum float32 = 123.23323423424242424
	var fnum2 float64 = 123.23323423424242424//for more precision
	fmt.Println("fnum is : ",fnum,",",fnum2)

	//IMPLICIT TYPE

	var st = "Happy"//but you cant change it lator on as lexer marks it stirng
	fmt.Println("string st is : ",st);

	// st = 10;//wrong

	//No var/wallerus operator

	x := 10;//but only allowed inside method 
	fmt.Println("val is : ",x)
}