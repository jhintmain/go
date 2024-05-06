package main

import "fmt"

/*
	1. bool (true, false)
	2. string ("Hello, World!") // 한번 생성되면 수정될 수 없는 Immutable 타입
	3. int  int8  int16  int32  int64 uint uint8 uint16 uint32 uint64 uintptr
	4. float32 float64 complex64 complex128	 // 복소수
	5. byte // alias for uint8 // 바이트 코드
	6. rune // alias for int32 // 유니코드 코드 포인트	
*/

func main(){
	// 문자 리터럴 : Back Quote(` `) 혹은 이중인용부호(" ")

	// Back Quote(` `) : Raw String Literal
	// 이스케이프 문자를 사용하지 않고 문자열 그대로 출력
	rawLiteral := `아리랑\n
아리랑\n
아라리요`
	interLiteral := "스리스리랑\n스리리요"
	interLiteral2 := "아리랑아리랑\n" +	"아리리요"	
	
	fmt.Println(rawLiteral)
	fmt.Println(interLiteral)
	fmt.Println(interLiteral2)

	// 타입 변환
	var i int = 100
    var u uint = uint(i)
    var f float32 = float32(i)  
    println(f, u)
  
    str := "ABC"
    bytes := []byte(str)
    str2 := string(bytes)
    println(bytes, str2)

	// 포인터 연산자
	var k int = 10
	var p = &k		//k의 주소를 할당
	println(p)		//p가 가리키는 주소를 출력
	println(*p)		//p가 가리키는 주소에 있는 실제 내용을 출력

}