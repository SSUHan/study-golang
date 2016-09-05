package main

import(
	"fmt"
)

// panic 함수는 문법적인 에러는 아닐지라도 로직에 따라 에러로 처리하고 싶은 것들을 처리한다
// recover 는 프로그램에 에러가 발생해 죽는 패닉이 일어났을때, 예외처리를 하기 위해서 사용한다
// recover 는 프로그램이 종료되기 전에 반드시 실행되어야 하므로 defer 키워드와 함께 사용이 되어야 한다
// 이 문법은 타 언어의 try ~ catch 를 구현하고 싶었던 것 같다.
// 어차피 recover 함수가 항상 마지막에 불려야 한다면, 내부적으로 recover 함수를 선언을 하는 것마으로도 defer 처리가 되도록
// 구현하면 좋지 않았을까 생각해본다 

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

