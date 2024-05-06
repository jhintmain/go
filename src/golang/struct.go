package main

import "fmt"

// struct 정의 - 내부
type person struct {
	name   string
	age    int
	gender string
}

// struct 정의 - 외부
type Person struct {
	name string
	age  int
}

func main() {
	// Go에는 전통적인 OOP 언어가 가지는 클래스, 객체, 상속 개념이 없다

	// 1. struct 정의
	p := person{}

	// struct 초기화 - 방법 1
	p.name = "Lee"
	p.age = 10
	p.gender = "M" // Fix: Changed 'M' to "M"

	fmt.Println(p)

	// struct 초기화 - 방법 2
	var p1 person
	p1 = person{"Kim", 20, "F"}
	p2 := person{"Park", 30, "M"}

	fmt.Println(p1)
	fmt.Println(p2)

	// struct 초기화 - 방법 3
	// new()를 사용하면 모든 필드를 Zero value로 초기화하고
	// person 객체의 포인터(*person)를 리턴

	p3 := new(person)
	p3.name = "Choi"
	p3.age = 40
	p3.gender = "F"

	fmt.Println(*p3)

	// 생성자 함수를 사용하여 초기화
	dic := newDict() // 생성자 호출
	dic.data[1] = "A"
	fmt.Println(*dic)

}

type dict struct {
	data map[int]string
}

func newDict() *dict {
	d := dict{}
	d.data = map[int]string{}
	return &d
}
