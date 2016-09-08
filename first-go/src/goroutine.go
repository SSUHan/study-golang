package main

import(
	"fmt"
	"math/rand"
	"time"
	"runtime"
)

func hello(n int){
	r := rand.Intn(100) // 랜덤 숫자 생성
	time.Sleep(time.Duration(r))
	fmt.Println(n)
}

func main(){
	// for i:=0; i<30; i++{
	// 	go hello(i)
	// }

	// Go 에서 멀티 코어를 활용하는 방법
	runtime.GOMAXPROCS(runtime.NumCPU()) // CPU 개수를 구한 뒤 사용할 최대 CPU 개수를 설정한다
	fmt.Println(runtime.GOMAXPROCS(0)) // 0을 넣으면 설정값을 그대로 출력해준다

	s := "Hello World"

	for i:=0;i<30;i++{
		go func(n int){
			fmt.Println(s, n)
		}(i) // 익명함수로 고루틴 실행
	}
	fmt.Scanln()
}