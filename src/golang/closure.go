package main

func nextValue() func()int{
	i:=0
	return func()int{
		i++
		return i
	}
}
func main(){
	next := nextValue()
	println(next())	// 	1
	println(next()) // 2
	println(next())	// 3

	anotherNext := nextValue()	// 새로운 클로저 생성
	println(anotherNext())	// 1	
	println(anotherNext())	// 2
}