package main

import (
	"fmt"
)

func main() {
	// Go 채널은 그 채널을 통하여 데이타를 주고 받는 통로
	// 채널은 make 함수를 사용하여 생성한다.
	// 채널은 <- 연산자를 사용하여 값을 주고 받는다.
	// 채널은 기본적으로 블록킹되어 있어서, 채널에 값을 주고 받을 때까지 대기한다. = 별도의 lock을 걸지 않고 데이타를 동기화하는데 사용
	// 채널은 Go루틴 간의 통신을 위해 사용된다.

	ch := make(chan int)

	go func() {
		ch <- 123 // 채널에 123을 보낸다
	}()

	println("111111")
	var i int
	i = <-ch // 채널로부터 값을 받는다
	// 채널로부터 데이타를 받고 있는데, 상대편 goroutine에서 데이타를 전송할 때까지는 계속 대기하게 된다
	// 채널로부터 데이타를 받으면, 그 때부터 다음 코드를 실행한다
	// 기다리는 코드 필요없음
	println("222222")
	println(i) // 123
	println("333333")

	// 기다리는 속성을 이용하여 아래 예제처럼 사용할 수도 있다
	done := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			println(i)
		}
		done <- true
	}()

	<-done //  위의 Go루틴이 끝날 때까지 대기

	// channel은 기본적으로 송/수신자가 있어야 한다, 없으면 데드락(에러)
	// 송신 : ch <- 1
	// 수신: <- ch

	// 채널 버퍼링
	// 채널 2가지 종류 : unbuffered channel, buffered channel
	// unbuffered channel : 채널에 버퍼가 없는 경우(위의 예제)
	// buffered channel : 채널에 버퍼가 있는 경우

	// Buffered Channel
	// 수신자가 받을 준비가 되어 있지 않을 지라도 지정된 버퍼만큼 데이타를 보내고 계속 다른 일을 수행할 수 있다
	// make(chan type, N) 함수를 통해 생성되는데, 두번째 파라미터 N에 사용할 버퍼 갯수를 넣는다
	ch1 := make(chan int, 1)

	//수신자가 없더라도 보낼 수 있다.
	ch1 <- 101

	fmt.Println(<-ch1) // print에서 찍는 수신은 수신으로 안침

	// 송/수신 전용 채널 함수 사용
	ch2 := make(chan string, 1)
	sendChan(ch2)
	receiveChan(ch2)

	// 채널 닫기
	// 채널 닫으면 송신X, 수신O
	// 수신은 2가지 return값을 가진다. (두번째 값은 채널이 닫혔는지 여부)
	ch3 := make(chan int, 2)

	// ch3에 1, 2를 보낸다
	ch3 <- 1
	ch3 <- 2

	// 채널 닫기
	close(ch3)

	// 채널 수신
	println(<-ch3)
	println(<-ch3)

	// if _, success := <-ch3; !success {
	// 	println("No more data")
	// }

	// 채널이 닫힌 것을 감지할 때까지 계속 수신
	for i := range ch3 {
		println(i)
	}

}

// 채널은 함수의 파라미터로도 사용할 수 있다
// 송신전용 채널
func sendChan(ch chan<- string) {
	ch <- "Data"
	// x := <-ch // 에러발생
}

// 수신전용 채널
func receiveChan(ch <-chan string) {
	data := <-ch
	fmt.Println(data)
}
