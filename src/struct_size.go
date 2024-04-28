package main

import (
	"fmt"
	"unsafe"
)

type struct1 struct{
	vBool bool // 1 byte
	Vint16 int16	//2 byte
	Vint32 int32	//4 byte
}

type struct2 struct{
	vBool bool 		//1 byte
	Vint16 int16	//2 byte
	Vint32 int32	//4 byte
	Vint64 int64	//8 byte
}

type struct3 struct{
	vBool bool 		//1 byte
	Vint16 int16	//2 byte
	Vint32 int32	//4 byte
	Vint64 int64	//8 byte
	Vint8 int8		//1 byte
}

type struct3_1 struct{
	vBool bool 		//1 byte
	Vint8 int8		//1 byte
	Vint16 int16	//2 byte
	Vint32 int32	//4 byte
	Vint64 int64	//8 byte	
}

func main(){
	fmt.Printf("Size of %T : %d byte\n", int64(0), unsafe.Sizeof(int64(0)))
	fmt.Printf("Size of %T : %d byte\n", int32(0), unsafe.Sizeof(int32(0)))
	fmt.Printf("Size of %T : %d byte\n", int16(0), unsafe.Sizeof(int16(0)))
	fmt.Printf("Size of %T : %d byte\n", int8(0), unsafe.Sizeof(int8(0)))
	fmt.Printf("Size of %T : %d byte\n", bool(false), unsafe.Sizeof(bool(false)))

	var s1 struct1
	fmt.Printf("Size of %T : %d byte\n", s1, unsafe.Sizeof(s1)) // 8 byte
	// 가장큰 byte(4) 기준으로 1+2+1(padding) + 4 = 8byte

	var s2 struct2
	fmt.Printf("Size of %T : %d byte\n", s2, unsafe.Sizeof(s2)) // 16 byte
	// 가장큰 byte(8) 기준으로 1+2+4+1(padding) + 8 = 16byte

	var s3 struct3
	fmt.Printf("Size of %T : %d byte\n", s3, unsafe.Sizeof(s3)) // 24 byte
	// 가장큰 byte(8) 기준으로 1+2+4+1(padding) + 8 + 1+7(padding)= 24byte

	var s3_1 struct3_1
	fmt.Printf("Size of %T : %d byte\n", s3_1, unsafe.Sizeof(s3_1)) // 16 byte
	// 가장큰 byte(8) 기준으로 1+1+2+4 + 8 = 16 byte

	
	/**
	type StringHeader struct {
		Data uintptr	// 8 byte
		Len  int		// 8 byte
	}
	**/
	var str string = "Hello"
	stringSize := len(str) + int(unsafe.Sizeof(str))
	fmt.Print(stringSize)	// 21 (문자열길이(5) + StringHeader 크기(8+8))


	type strInclude struct {
		vStr	string // StringHeader(8+8) byte
		vInt64	int64	// 8 byte
	}
	...
	fmt.Println(unsafe.Sizeof(strInclude{}))// 24 byte
	// 가장큰 byte(8) >  8 + 8 + 8 = 24 

}