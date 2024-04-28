package main

import "fmt"

func main(){
	var name string
	fmt.Print("name : ")
	fmt.Scanf("%s", &name)

	switch name {
		case "vamos":
			fmt.Printf("name is %s\n",name)
			fmt.Println("but it is not a name")
		case "jhpark":
			fmt.Printf("name is %s\n",name)
			fmt.Println("yes it it your name")
		default : 
			fmt.Println("other..")
		
	}
}