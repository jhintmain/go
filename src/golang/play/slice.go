package main

import "fmt"

func main(){
	// slice 선언 
	// 방법 1. 배열을 선언하듯이 "var v []T" 처럼 하는데 배열과 달리 크기는 지정하지 않는다.
	var a []int        //슬라이스 변수 선언
    a = []int{1, 2, 3} //슬라이스에 리터럴값 지정
    a[1] = 10
    fmt.Println(a)     // [1, 10, 3]출력

	// 방법 2. 내장함수 make() 함수를 이용할 수 있다. make() 함수로 슬라이스를 생성하면, 개발자가 슬라이스의 길이(Length)와 용량(Capacity)을 임의로 지정할 수 있는 장점이 있다
	s := make([]int, 5, 10)
    fmt.Println(len(s), cap(s)) // len 5, cap 10


	// nil 슬라이스
	var m []int
 
    if m == nil {
        println("Nil Slice")
    }
    println(len(m), cap(m)) // 모두 0

	// 부분 슬라이드 : [처음인덱스:마지막인덱스+1]
	c := []int{0, 1, 2, 3, 4, 5}
    c = c[2:5]  
    fmt.Println(c) //2,3,4 출력

	// 슬라이스 인덱스는 처음/마지막 둘 중 하나 혹은 둘 다를 생략할 수도 있다. 처음 인덱스가 생략되면 0 이, 마지막 인덱스가 생략되면 그 슬라이스의 마지막 인덱스가 자동 대입된다. 
	// 즉, 처음 0 부터 인덱스 4까지를 포함하기 위해서는 [:5]를,
	// 인덱스 2부터 마지막까지 포함하기 위해서는 [2:]를 사용한다.
	// 만약 [:]와 같이 모두 생략하면, 전체를 표현한다.
	t := []int{0, 1, 2, 3, 4, 5}
	t = t[2:5]     // 2, 3, 4
	t = t[1:]      // 3, 4
	fmt.Println(t) // 3, 4 출력

	// 배열은 고정된 크기로 그 크기 이상의 데이타를 임의로 추가할 수 없지만, 슬라이스는 자유롭게 새로운 요소를 추가할 수 있다.

	// 내장함수인 append() 함수를 이용하여 슬라이스에 새로운 요소를 추가할 수 있다.
	// append() 함수는 슬라이스에 새로운 요소를 추가한 후, 그 결과를 반환한다.
	// 만약 슬라이스의 용량(capacity)을 초과하여 요소를 추가하면, 슬라이스의 용량(capacity)은 자동으로 늘어난다.
	  
	b := append(t, 7,8,9) // 3,4,7,8,9
	fmt.Println(b)

	sliceA := make([]int, 0, 3)
 
    // 계속 한 요소씩 추가
    for i := 1; i <= 15; i++ {
        sliceA = append(sliceA, i)
        // 슬라이스 길이와 용량 확인
        fmt.Println(len(sliceA), cap(sliceA))
    } 
    fmt.Println(sliceA) // 1 부터 15 까지 숫자 출력 

	sliceB := []int{1, 2, 3}
    sliceC := []int{4, 5, 6}
 
    sliceB = append(sliceB, sliceC...)
    //sliceB = append(sliceB, 4, 5, 6)
 
    fmt.Println(sliceB) // [1 2 3 4 5 6] 출력

}