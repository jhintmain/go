package main

import(
	"fmt"
	"sync"
)

type pTest struct{
	pStr string
	pInt int
}

func addStr(wg *sync.WaitGroup, getStruct *pTest, str rune){
	getStruct.pStr = getStruct.pStr +string(97+str)
	defer wg.Done()
}

func addInt(wg *sync.WaitGroup , getStruct *pTest, num int){
	getStruct.pInt += num
	defer wg.Done()
}

func printStruct(wg *sync.WaitGroup, getStruct *pTest, num int){
	fmt.Println(getStruct)
	defer wg.Done()
}

func main(){
	var example pTest
	defer fmt.Println("====== Finished =====")
	wg := sync.WaitGroup{}
	defer wg.Wait()

	for i :=0; i<=10; i++{
		wg.Add(1)
		go addStr(&wg, &example, rune(i))
		wg.Add(1)
		go addInt(&wg, &example, i)
		wg.Add(1)
		go printStruct(&wg, &example, i)
	}
}