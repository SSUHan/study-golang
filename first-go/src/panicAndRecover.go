package main

import(
	"fmt"
)

func f(){
	defer func(){
		r := recover()
		fmt.Println(r)
	}()

	a := [...]int{1,2} // 길이가 2 인 배열

	for i := 0; i<5; i++{
		// 강제로 인덱스가 5까지 돌려서 out of bound 에러를 내보자
		fmt.Println(a[i])
	}
}

func main(){
	f()
	fmt.Println("after panic")
}

