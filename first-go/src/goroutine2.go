package main 

import (
	"fmt"
	"runtime"
)

func main(){
	runtime.GOMAXPROCS(1) // Single Core 로 동작하도록 함

	s:="Hello World"
	fmt.Println("===인자로 복사해놓고 go하는경우===")
	for i:= 0; i<30; i++{
		go func(n int){ // go routine 이 돌 함수의 인자값으로 복사해서 넘겨 받아 진행하는 경우
			fmt.Println(s, n) 
		}(i)
	}
	fmt.Println("=========================")
	fmt.Scanln()
	fmt.Println("===인자로 안받고 외부변수를 참조해서 go하는경우===")
	for i:= 0; i<30; i++{
		go func(){ // go routine 이 돌 함수가 for 문의 변수를 참조하는 경우
			fmt.Println(s, i)
		}()
	}
	fmt.Println("=========================")
	fmt.Scanln()

	/*
		클로저를 고루틴으로 실행할때 반복문에 의해 바뀌는 변수는 반드시 매개변수로 넘겨준다
		매개변수로 넘겨준 시점에서 해당 변수의 값이 복사가 되기 때문에, 고루틴에서 값의 변화를 걱정할 필요가 없다
		따라서, CPU 코어를 하나만 사용하던, 여러개를 사용하던 상관없이 반복문에 의해 바뀌는 변수는 매개변수로 넘겨주어야 한다
	*/

}