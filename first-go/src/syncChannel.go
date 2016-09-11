package main 

import "fmt"
import "time"

// 채널 선언 방법 make(chan <자료형>)
func sum(a int, b int, c chan int){
	c <- a + b // 채널 c에 a+b을 보낸다 
}

func main(){
	c := make(chan int) // int 를 담는 채널을 동적 할당 한다

	go sum(1, 2, c)

	n := <-c 		// 채널에서 값을 꺼내, n 에 대입
	fmt.Println(n)

	done := make(chan bool) // 인자를 채널의 타입만 집어넣었다. 이를 동기채널이라고 하며, 버퍼크기는 1개이다.
	count := 3

	go func(){
		for i:=0; i<count; i++{
			done <- true				// 하나를 인풋하고, 버퍼가 1개일뿐이므로, 다음 아이템을 집어넣지 못하고 기다리게 된다.
			fmt.Println("고루틴 : ", i)
			time.Sleep(1*time.Second)
		}
	}()

	for i:=0; i<count; i++{
		<-done
		fmt.Println("메인함수 : ", i)
	}
}