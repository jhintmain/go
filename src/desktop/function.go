package main
import "fmt"


func NameOfFunc(a int , b int ) (isTrue bool){
	isTrue = (a == b )
	return 
}

func nameOfFunc(a,b int) bool{
	isTrue :=(a==b)
	return isTrue;
}

func main(){
	var rs bool = NameOfFunc(3,3)
	var rs2 bool = nameOfFunc(2,3)

	fmt.Print(rs,rs2)
}