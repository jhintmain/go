package main

import "fmt"

func main(){
	// if문 : 조건은 항상 Boolean
	var k int = 1
	if k == 1 {
		fmt.Println("One")
	} else if k == 2 {
		fmt.Println("Two")
	} else {
		fmt.Println("Other")
	}

	// switch문 : break문이 필요 없음 1개의 case만 실행
	switch k {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3,4:
		fmt.Println("Three or Four")
	default:
		fmt.Println("Other")
	}

	grade(95)

	check(2)

	typeSwitch(1.0)
}

// switch문 : switch 뒤에 조건변수 혹은 Expression을 적지 않는 용법이다.
func grade(score int) {
	switch {
	case score >= 90:
		println("A")
	case score >= 80:
		println("B")
	case score >= 70:
		println("C")
	case score >= 60:
		println("D")
	default:
		println("No Hope")
	}
}   

// switch문 : fallthrough (다음 case문을 실행)
func check(val int) {
    switch val {
    case 1:
        fmt.Println("1 이하")
        fallthrough
    case 2:
        fmt.Println("2 이하")
        fallthrough
    case 3:
        fmt.Println("3 이하")
        fallthrough
    default:
        fmt.Println("default 도달")
    }
}

func typeSwitch(v interface{}) {
	// switch문 : Type Switch (변수 v의 타입이 int 인지, bool, string 인지를 체크)
	switch v.(type) {
	case int:
		println("int")
	case float64:
		println("float64")
	case bool:
		println("bool")
	case string:
		println("string")
	default:
		println("unknown")
	}   
}