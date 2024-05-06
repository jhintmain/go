package main

import (
	"fmt"
	"os"
)

func main() {
	// defer()
	// : 특정 문장 혹은 함수를 나중에 (defer를 호출하는 함수가 리턴하기 직전에) 실행하게 한다
	// defer는 LIFO (Last In First Out) 순서로 실행된다.
	// defer는 함수의 실행이 완료된 후에 실행된다.
	// 일반적으로 defer는 finally 블럭처럼 마지막에 Clean-up 작업을 위해 사용된다.

	f, err := os.Open("tt.txt")
	if err != nil {
		panic(err)
	}

	// main 마지막에 파일 close 실행
	defer f.Close()

	// 파일 읽기
	bytes := make([]byte, 1024)
	f.Read(bytes)
	println(len(bytes)) // 1024

	// panic()
	// : 프로그램이 비정상적으로 종료될 때 사용한다.
	// panic 함수는 현재 함수를 즉시 멈추고
	// 현재 함수의 defer 함수들을 모두 실행한 후 즉시 리턴한다.
	// 마지막에는 프로그램이 에러 메시지와 함께 종료된다.

	// 이러한 방식을 통해 panic은 상위 함수에도 영향을 미친다.
	// 그리고 상위 함수에서도 계속 콜스택을 타고 올라가며 defer 함수들을 실행한다.

	// 잘못된 파일명을 넣음
	openFile("Invalid.txt")

	// openFile() 안에서 panic이 실행되면
	// 아래 println 문장은 실행 안됨
	println("Done") // 실행 안됨

	// recover()
	// : panic 함수에 의한 패닉상태를 다시 정상상태로 되돌리는 함수이다.
}

func openFile(fn string) {
	// defer 함수. panic 호출시 실행됨
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("OPEN ERROR", r)
		}
	}()

	f, err := os.Open(fn)
	if err != nil {
		panic(err) // panic 발생 > 현재 함수의 defer들 실행
	}

	defer f.Close() // defer 실행
}
