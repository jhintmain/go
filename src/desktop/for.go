package main

import "fmt"

func main(){

	for i :=0; i<100; i++{
		fmt.Printf("%d ",i)
	}

	fmt.Printf("\n----------------\n")

	j :=0
	for j  < 100{
		fmt.Printf("%d ",j)
		j++;
		if j > 100{
			break
		}
	}

	fmt.Printf("\n----------------\n")
	var rgData [4]string
	rgData[0] = "Hi"
	rgData[1] = "My"
	rgData[2] = "Name"
	rgData[3] = "Jihea Park"

	for i,v := range rgData{
		fmt.Printf("index : %d \t value : %s \n",i,v)
	}

	fmt.Printf("\n----------------\n")
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	
	a[2] = 10
	for i :=0; i<len(a); i++{
		fmt.Printf("%d ",a[i])
	}

	// fmt.Printf("\n----------------\n")
	// i := 0
	// for i <len(a){
	// 	fmt.Printf("%d",a[i])
	// }

	fmt.Printf("\n----------------\n")

	for i, v := range a{
		fmt.Printf("index %d & value : %d \n",i,v);
	}
}