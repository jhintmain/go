package main

import "fmt"

func main() {	
	var x int8 = 5
	var y int8 = 3
	fmt.Print("x is ", x, "\n", "y", "is", y, "\n")

	var a string
	var b string
	fmt.Scan(&a, &b) // string 두 개의 값을 입력 받아, 각각 x와 y에 저장합니다.
	// Hello World!! 
	// 또는
	// Hello
	// World!!
	// z, _ := fmt.Scan(&x, &y) // 인자의 개수를 z에 초기화합니다.
	fmt.Println(a, "and", b)
	// Hello and World!!
	// fmt.Println(z)
}
