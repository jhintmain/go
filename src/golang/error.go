package main

import (
	"log"
	"os"
)

func main() {
	// Go 언어는 예외 처리를 위한 try-catch-finally 문법을 제공하지 않는다.
	// 대신 Go 언어는 에러를 반환하는 방식으로 예외 처리를 한다.
	// Go 언어에서는 함수가 에러를 반환할 수 있도록 하기 위해, 함수의 반환값으로 error 타입을 반환하도록 한다.
	// error 타입은 Go 언어의 내장 타입으로, error 인터페이스를 구현하고 있다.
	// error 인터페이스는 Error() 메서드를 가지고 있으며, 이 메서드는 에러 메시지를 반환한다.
	// 에러가 발생하지 않았다면, nil 값을 반환한다.

	f, err := os.Open("tt.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	println(f.Name())

}
