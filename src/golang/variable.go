package main

import "fmt"

// # Go 에서는 할당되는 값을 보고 그 타입을 추론하는 기능이 자주 사용된다

func main(){
	// =====  변수 =====

	// 1. var [이름] [타입] = [값](option)
	var a int
	var f float32 = 11

	a = 10
	fmt.Println(a, f)

	// 2. var [이름, 이름, 이름] [타입] = [값, 값, 값](option)
	var i,j,k int = 1,2,3
	var c, python, java string

	c="c"
	python="python"
	java="java"

	fmt.Println(i, j, k, c, python, java)

	// 3. var [이름] = [값] === [이름] := [값] > 함수내에서만 사용 가능
	var x = 10
	y := 20

	fmt.Println(x, y)


	// ===== 상수 =====

	// 1. const [이름] [타입] = [값]
	const d int = 10
	const s string = "Hi"
	fmt.Println(d, s)

	// 2. 여러 개의 상수들 묶어서 지정
	const (
		Visa = "Visa"
		Master = "MasterCard"
		Amex = "American Express"
	)
	fmt.Println(Visa, Master, Amex)

	// 3. iota : 상수값 생성기 (iota는 0부터 시작 1씩 증가)
	const (
		Apple = iota // 0
		Grape        // 1
		Orange       // 2
	)
	fmt.Println(Apple, Grape, Orange)
}