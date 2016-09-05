package main

import(
	"fmt"
)

// 지연 호출은 특정 함수를 현재 함수가 끝나기 직전에 실행시켜달라는 기능
// finally 의 기능과 비슷

func hello(){
	fmt.Println("hello")
}

func world(){
	fmt.Println("world")
}

func main(){
	defer world()
	hello()
	hello()
	hello()

	helloworld()

	// 지연 함수 호출되는 순서는 스택과 동일
	// 가장 나중에 호출한 지연함수가 가장 먼저 호출되도록 구조화되어 있다.
	// 지연 함수는 파일 I/O 에서 유용하게 사용된다고 한다
}

func helloworld(){
	defer func(){
		fmt.Println("world2")
	}() // 익명함수로 실행 하지만 defer 걸어둠

	func(){
		fmt.Println("hello2")
	}() // 익염함수로 바로 실행
}