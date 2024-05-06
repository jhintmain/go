package main

// import "fmt"

func main(){
	// Go에서 파라미터를 전달하는 방식은 크게 Pass By Value와 Pass By Reference로 나뉜다.
	var msg string = "Hello"

	// 1. Pass By Value : 값에 의한 전달
	// - 함수에 인자를 전달할 때, 값이 복사되어 전달된다.
	sayByvalue(msg)

	// 2. Pass By Reference : 참조에 의한 전달
	// - 함수에 인자를 전달할 때, 값이 복사되지 않고 참조만 전달된다.
	sayByreference(&msg)
	println(msg)

	//  Variadic Function (가변인자함수)
	// - 함수의 파라미터를 가변적으로 받을 수 있는 함수
	// - 함수의 파라미터를 정확히 모를 때 사용
	sayByVariadic("This", "is", "a", "book")

	// Go 언어는 복수개의 값을 리턴할 수 있다.
	// Named Return Parameter 라는 기능을 제공하는데, 이는 리턴되는 값들을 (함수에 정의된) 리턴 파라미터들에 할당할 수 있다.
	count, total := sum1(1, 7, 3, 5, 9)
    println(count, total) 

	count, total = sum2(1, 7, 3, 5, 9, 11)
	println(count, total) 

	// 익명함수 지원
	sum := func(n ...int) int { //익명함수 정의
        s := 0
        for _, i := range n {
            s += i
        }
        return s
    }
 
    result := sum(1, 2, 3, 4, 5) //익명함수 호출
    println(result)

	// 일급함수 지원
	// - 함수를 변수에 대입하거나, 다른 함수의 파라미터로 전달하거나, 다른 함수의 리턴값으로 사용할 수 있는 함수
    
	// 변수 add 에 익명함수 할당
    add := func(i int, j int) int {
        return i + j
    }
 
    // add 함수 전달
    r1 := calc(add, 10, 20)	// add 함수를 파라미터로 전달
    println(r1)
 
    // 직접 첫번째 파라미터에 익명함수를 정의함
    r2 := calc(func(x int, y int) int { return x - y }, 10, 20)
    println(r2)	

}
// Delegate -원형 정의
type calculator func(int, int) int
 
// Delegate (o) - calculator 원형 사용
func calcDelegate(f calculator, a int, b int) int {
    result := f(a, b)
    return result
}
// Delegate (x)- calculator 원형 사용
func calc(f func(int, int) int, a int, b int ) int {
	result := f(a, b)
	return result
}

func sayByvalue(msg string) {
    println(msg)
}

func sayByreference(msg *string) {
    println(*msg)
    *msg = "Changed" //메시지 변경
}

func sayByVariadic(msg ...string) {
	for _, v := range msg {
		println(v)
	}
}

func sum1(nums ...int) (int, int) {
	sum := 0	// 합계
	cnt := 0  // 요소 갯수
    for _, n := range nums {
        sum += n
        cnt++
    }
    return cnt, sum
}

func sum2(nums ...int) (count int, total int) {
    for _, n := range nums {
        total += n
    }
    count = len(nums)
    return // return count, total
}