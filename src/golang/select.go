package main

import (
	"time"
)

func main() {

	done1 := make(chan bool)
	done2 := make(chan bool)

	go run1(done1)
	go run2(done2)

	// select 문은 여러 채널들을 기다리면서 준비된 채널이 있으면 실행하는 구문이다.

EXIT: // 레이블
	for { // 무한 : for문을 무한으로 돌리고, select문을 통해 채널을 기다린다.
		select {
		case <-done1:
			println("run1 완료")

		case <-done2:
			println("run2 완료")
			break EXIT // 레이블을 통해 for문을 빠져나온다.
		}
	}
}

func run1(done1 chan bool) {
	time.Sleep(1 * time.Second)
	done1 <- true
}

func run2(done2 chan bool) {
	time.Sleep(2 * time.Second)
	done2 <- true
}
