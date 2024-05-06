package main

import (
	"fmt"
)

type Shape interface {
	Area() float64
}

type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}

// Shape 인터페이스를 구현하기 위해 Area() 메서드를 구현한다.
func (r Rect) Area() float64 {
	return r.width * r.height
}

// Shape 인터페이스를 구현하기 위해 Area() 메서드를 구현한다.
func (c Circle) Area() float64 {
	return c.radius * c.radius * 3.14
}

// interface{} 는 모든 타입을 나타내는 빈 인터페이스이다.
func printValue(value interface{}) {
	fmt.Println("Value:", value)
}

func main() {
	// 구조체(struct)가 필드들의 집합체라면, interface는 메서드들의 집합체이다.
	//  interface는 타입(type)이 구현해야 하는 메서드 원형(prototype)들을 정의한다

	shapes := []Shape{
		Rect{10, 20}, // Shape 의 Area() 메서드를 구현하여 Shape type으로 취급 가능
		Circle{15},   // Shape 의 Area() 메서드를 구현하여 Shape type으로 취급 가능
	}

	for _, shape := range shapes {
		fmt.Println(shape.Area())
	}

	// Type Assertion
	var x interface{} = 1
	i := x       // type: interface{}, value: 1
	j := x.(int) // type: int, value: 1

	fmt.Println(i) // 1
	fmt.Println(j) // 1

	printValue(1)
	printValue("Hello")
	printValue(3.14)
}
