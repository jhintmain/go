package main

import "fmt"


func addFunc(a,b int) int {
	return a+b
}

func subFunc(a,b int) int {
	return a-b
}

func mulFunc(a,b int) int {
	return a*b
}
func divFunc(a,b int) int {
	return a/b
}
func modFunc(a,b int) int {
	return a%b
}
func main(){
	var x,y int
	fmt.Println("계산식을 위한 2개의 값을 입력하세요")
	fmt.Println("foramt : value1 value2 then press Enter")
	fmt.Scanf("%d %d ", &x, &y)
	fmt.Println("---------------------------")
	fmt.Println("result of addFunc(x,y) : ",addFunc(x,y))
	fmt.Println("result of subFunc(x,y) : ",subFunc(x,y))
	fmt.Println("result of mulFunc(x,y) : ",mulFunc(x,y))
	fmt.Println("result of divFunc(x,y) : ",divFunc(x,y))
	fmt.Println("result of modFunc(x,y) : ",modFunc(x,y))
}