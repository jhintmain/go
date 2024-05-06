package main

import "fmt"

type Rect struct {
	width, height int
}

// 메소드 선언 - Value Receiver
func (r Rect) area() int {
	return r.width * r.height
}

// 메소드 선언 - Pointer Receiver
func (r *Rect) area2() int {
	r.width++
	return r.width * r.height
}

func main() {
	// Go에서 메서드는 특별한 형태의 함수이다.
	// 메서드는 함수 정의에서 func 키워드와 함수명 사이에 "그 함수가 어떤 struct를 위한 메서드인지"를 표시하게 된다

	rect := Rect{10, 20}
	area := rect.area()     // 메소드 호출
	fmt.Println(rect, area) // {10,20} 200

	area2 := rect.area2()    // 메소드 호출
	fmt.Println(rect, area2) // {11,20} 220
}
