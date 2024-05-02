package main

import (
	"fmt"
	"math/rand"
)

func main(){

	target := rand.Intn(100) + 1	// 0~99 +1 => 1~100 사이 난수
	var input int
	var success bool = false

	for i:=0; i<10; i++ {
		fmt.Print("Enter your guess : ")
		fmt.Scanf("%d",&input)
		if input > target{
			fmt.Println("Oops.Your guess was HIGH")
		}else if input < target{
			fmt.Println("Oops.Your guess was LOW")	
		}else{
			success = true
			fmt.Println("Good job! You guessed it!")	
			break
		}
	}

	if success == false{
		fmt.Println("Sorry you didn't guess my number. It was Number : ",target)
	}
}