package main

import "fmt"

func main() {
	// Go 언어는 Map 타입을 내장하고 있는데, "map[Key타입]Value타입" 과 같이 선언할 수 있다
	// map은 reference 타입이므로, zerovalue로  nil 값을 갖는다.

	var idMap map[int]string // map 선언 (nil map)

	// 초기화 방법 1. make 사용
	idMap = make(map[int]string) // map을 초기화하기 위해 make()함수를 사용할 수 있다

	// add or update
	idMap[901] = "Apple"
	idMap[134] = "Grape"
	idMap[777] = "Tomato"

	str := idMap[134]
	println(str) // Grape 출력

	noData := idMap[999] // 값이 없으면 value type의 zero-value 반환
	println(noData)      // string의 zero-value "" 출력

	println(idMap[777])
	delete(idMap, 777) // 삭제
	println(idMap[777])

	// 초기화 방법 2. 리터럴을 사용한 초기화
	tickers := map[string]string{
		"GOOG": "Google Inc",
		"MSFT": "Microsoft",
		"FB":   "FaceBook",
	}

	println(tickers["GOOG"])

	// map key 체크하기
	// map 키 체크
	val, exists := tickers["MSFT"] // 2개의 리턴값을 리턴 ,  첫번째는 키에 상응하는 값이고, 두번째는 그 키가 존재하는지 아닌지를 나타내는 bool 값이다.
	if !exists {
		println("No MSFT ticker")
	} else {
		println(val)
	}

	// map 순회
	// for range 문을 사용하여 모든 맵 요소 출력
	// **** Map은 unordered 이므로 순서는 무작위 ****
	for key, val := range tickers {
		fmt.Println(key, val)
	}
}
