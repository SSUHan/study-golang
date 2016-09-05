package main

import(
	"fmt"
)
// 클로저를 사용하면 지역변수가 소멸되지 않고,
// 나중에 함수를 호출할 때마다 계속 가져다 쓸 수 있다.
// 즉, 클로저는 함수가 선언될 떼의 환경을 계속 유지한다.
func calc() func(x int) int{
	a, b := 3, 5
	return func(x int) int{
		return a*x + b 
	}
}

func main(){
	f := calc()

	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
	fmt.Println(f(4))
}