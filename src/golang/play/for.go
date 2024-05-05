package main

import "fmt"

func main(){
	// for문
	// for 초기화; 조건; 후처리 {}
	// 초기화 : 조건문이 참일때만 실행
	// 조건 : 반복문이 실행되는 조건
	// 후처리 : 조건이 거짓일때 실행
	for i := 0; i < 5; i++ {
		fmt.Print(i)
	}

	fmt.Println()
	// for문 : 초기화, 후처리 생략
	// 조건만 존재하면 while문과 같은 역할
	i := 0
	for i < 5 {
		fmt.Print(i)
		i++
	}
	fmt.Println()

	// break, continue, goto
	var a = 1
    for a < 15 {
        if a == 5 {
            a += a
            continue // for루프 시작으로
        }
        a++
        if a > 10 {
            break  //루프 빠져나옴
        }
    }
    if a == 11 {
        goto END //goto 사용예
    }
    println(a)
 
END:
    println("End")


	j := 0
 
	L1:
		for {
		 
			if j == 0 {
				break L1
			}
		}
	 
		println("j break - OK")

	// for문 : 무한루프
	// for {
	// 	fmt.Println("Infinite Loop")
	// }

	// for문 : range
	// 배열, 슬라이스, 맵 등 순회할때 사용
	arr := [3]int{1, 2, 3}
	for i, v := range arr {
		fmt.Println(i, v) // index, value
	}

	// for문 : range
	// 문자열 순회
	str := "Hello, World!"
	for i, v := range str {
		fmt.Println(i, string(v))	// index, value
	}

	// for문 : range
	// 맵 순회
	m := map[string]int{"A": 10, "B": 20}
	for k, v := range m {
		fmt.Println(k, v)	// key, value
	}
}