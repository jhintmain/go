package main

func main(){
	// Go에서 배열의 배열크기는 Type을 구성하는 한 요소
	// 고정된 배열크기 안에 동일한 타입의 데이타를 연속적으로 저장
	// [3]int와 [5]int는 서로 다른 타입으로 인식
	var arr1 [3]int
	var arr2 [5]int
	println(arr1[0])
	println(arr2[0])
	
	var arr1_init = [3]int{1,2,3}
	var arr2_init = [...]int{1,2,3,4,5} // 배열크기 자동으로 할당
	println(arr1_init)
	println(arr2_init)

	// 다차원 배열
	var a = [2][3]int{
        {1, 2, 3},
        {4, 5, 6},  //끝에 콤마 추가
    }
    println(a[1][2])
}