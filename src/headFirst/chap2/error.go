package main

import (
	"fmt"
	"log"
	"os"
)

func main(){
	fileinfo,err := os.Stat("my1.txt")

	if err != nil{
		log.Fatal(err)		
	}
	
	fmt.Println(fileinfo.Size())	
}