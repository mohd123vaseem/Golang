package main

import "fmt"

func main() {

	nums := []string{"vaz", "naz", "vaseem", "kaz"}

	for i := 0; i < len(nums); i++ {
        
		if i==2 {
			// break;
			continue
		}else if i==len(nums)-1 {

			fmt.Println("nums ends here")
		}
		fmt.Println(nums[i],",")
	}

	for i := range nums {
         
		fmt.Println(nums[i])
	}

	for i,st := range nums {
         
		fmt.Printf("index : %v , string is : %s \n",i,st)
	}
}