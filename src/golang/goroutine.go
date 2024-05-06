package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// Go루틴(goroutine)
	// : Go 런타임이 관리하는 Lightweight 논리적 (혹은 가상적) 쓰레드(주1)이다.
	// "go" 키워드를 사용하여 함수를 호출하면, 런타임시 새로운 goroutine을 실행한다.
	// goroutine은 비동기적으로(asynchronously) 함수루틴을 실행하므로, 여러 코드를 동시에(Concurrently) 실행하는데 사용된다.
	// Go 런타임은 Go루틴을 관리하면서 Go 채널을 통해 Go루틴 간의 통신을 쉽게 할 수 있도록 하였다.
	// Go루틴은 쓰레드보다 적은 메모리를 사용하고, 쓰레드보다 빠르게 시작되며, 쓰레드보다 쉽게 사용할 수 있다.

	/*
		// 동기 실행
		say("Sync Call")

		// 비동기 실행
		go say("Async Call1")
		go say("Async Call2")
		go say("Async Call3")

		time.Sleep(time.Second * 3) // 3초 대기
	*/

	// WaitGroup 생성. 2개의 Go루틴을 기다림.
	var wait sync.WaitGroup
	wait.Add(2) // 몇개의 goroutine을 기다릴지 설정

	// 익명함수를 사용한 goroutine
	go func() {
		defer wait.Done() //끝나면 .Done() 호출
		fmt.Println("Hello")
	}()

	go func(msg string) {
		defer wait.Done() //끝나면 .Done() 호출
		fmt.Println(msg)
	}("Hi") //함수를 호출할 때 파라미터를 넘겨줌

	wait.Wait() //Go루틴 모두 끝날 때까지 대기

	// 다중 CPU 사용
	// : Go 런타임은 기본적으로 하나의 CPU만 사용하도록 설정되어 있다.
	// 하지만, GOMAXPROCS() 함수를 사용하여 사용할 CPU의 개수를 설정할 수 있다.
	// 시스템의 CPU 코어 수에 따라 자동으로 설정되며, 만약 설정을 변경하고 싶다면 GOMAXPROCS() 함수를 사용하면 된다.
	// runtime.GOMAXPROCS(runtime.NumCPU()) // CPU 개수 반환 후 설정
	fmt.Println(runtime.NumCPU())      // CPU 개수 반환
	fmt.Println(runtime.GOMAXPROCS(0)) // 설정값 반환

}

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s, "****", i)
	}
}
