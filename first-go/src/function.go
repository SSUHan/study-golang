package main

import(
	"fmt"
)



func main(){
	hello()
	fmt.Println(add(2,3))
	fmt.Println(addAndSub(5, 3))
	fmt.Println(sum(3,4,5))
	fmt.Println(sum(3,4,5,6,7,8,9,9))

	// 슬라이스 자체를 넘기고 싶을때는
	n := []int{1,2,3,4,5}
	r := sum(n...) // 이렇게 ... 를 붙이면 슬라이스의 요소 하나하나가 넘어가게 된다
	fmt.Println(r)

	// 함수를 변수에 저장할 수 있다.
	var h func(a int, b int) int = add
	h2 := add // 빠른 대응

	fmt.Println(h(1,2))
	fmt.Println(h2(3,4))

	// 함수를 슬라이스와 맵에도 저장할 수 있다.
	f_slice := []func(int, int)int{f1, f2}
	fmt.Println(f_slice[0](1,2))
	fmt.Println(f_slice[1](2,3))

	f_map := map[string]func(int, int)int{
		"sum":f1,
		"diff":f2,
	}
	fmt.Println(f_map["sum"](1,2))
	fmt.Println(f_map["diff"](3,5))

	// 익명함수를 정의하고 사용할 수 있다
	func(){
		fmt.Println("anomi func")
	}() // 실행까지 바로

	func(s string){
		fmt.Println(s)
	}("anomi hello world")


}

// go 는 함수 선언의 위치의 제약이 없다
func hello(){
	fmt.Println("Hello")
}

// 매개변수와 리턴값은 이렇게 정의한다
func add(a int, b int) int{
	return a+b
}

// 괄호 안에 리턴값 변수의 이름과 자료형을 지정하면,
// 함수 안에서 리턴값 변수를 사용할 수 있다.
// 함수 가장 마지막에 리턴을 사용한다
func add2(a int, b int) (r int){
	r = a + b
	return
}

// 리턴값을 여러개 반환할 수 있다.
// 사용하고 싶지 않은 리턴값은 _(밑줄문자) 를 사용합니다
func addAndSub(a int, b int) (int, int) {
	return a+b, a-b
}

// 가변인자 사용하기
// for ~ range 와 함께 사용하는 템플릿
// 가변인자로 받아오는 변수는 슬라이스 타입이므로 range 키워드로 꺼내서 사용할 수 있다
func sum(n ...int) int{
	total := 0
	for _, value := range n{
		total += value
	}
	return total
}

func f1(a int, b int)int{
	return a+b
}
func f2(a int, b int)int{
	return a-b
}