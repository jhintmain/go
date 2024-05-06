package main

import (
	"fmt"

	"jhpark.com/go/src/testlib" // alias 사용가능
)

func main() {
	// 일반적으로 패키지는 라이브러리로서 사용되지만,
	// "main" 이라고 명명된 패키지는 Go Compiler에 의해 특별하게 인식된다.
	// 패키지명이 main 인 경우, 컴파일러는 해당 패키지를 공유 라이브러리가 아닌 실행(executable) 프로그램으로 만든다.
	// 그리고 이 main 패키지 안의 main() 함수가 프로그램의 시작점, 즉 Entry Point가 된다.
	// 패키지를 공유 라이브러리로 만들 때에는, main 패키지나 main 함수를 사용해서는 안된다.

	// 다른 패키지를 프로그램에서 사용하기 위해서는 import 를 사용하여 패키지를 포함시킨다
	fmt.Println("Hello")

	// 패키지를 import 할 때, Go 컴파일러는 GOROOT 혹은 GOPATH 환경변수를 검색하는데, 표준 패키지는 GOROOT 안의 패키지에서 그리고 사용자 패키지나 3rd Party 패키지는 GOPATH/pkg 에서 패키지를 찾게 된다.

	fmt.Println(testlib.GetMusic("Adele"))

}
