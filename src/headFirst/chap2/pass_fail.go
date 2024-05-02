package main

import (
	"fmt"
)

func main() {
	var grade float64
	var status string
	
	fmt.Print("Enter a grade:")
	fmt.Scanf("%f", &grade)

	if grade >= 60{
		status = "passing"		
	}else{
		status = "failing"
	}

	fmt.Println("A grade of",grade,"is",status)
}