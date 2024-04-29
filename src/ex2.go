// package main

// import (
// 	"fmt"
// 	"unsafe"
// )

// type pTest struct{
// 	pStr string	// 16byte
// 	pInt int	// 8byte (64os)
// }

// func setStr(getStruct *pTest){
// 	getStruct .pStr = "struct pTest's byte size : "
// }

// func setInt(getStruct *pTest){
// 	getStruct.pInt = int(unsafe.Sizeof(pTest{}))
// }

// func main(){
// 	var example pTest

// 	setStr(&example)
// 	setInt(&example)

// 	fmt.Println(example)
// 	// {struct pTest's byte size :  24}
// }