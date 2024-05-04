package main

import "fmt"

func main(){
	// cannot use b (variable of type int64) as int value in assignment
	// var a int = 2
	// var b int64 = 3
	// a = b 


	var a [5]int = [5]int{1,2,3,4,5}
	var b [5]int = [5]int{6,7,8,9,10}
	a = b
	fmt.Println(a)
}