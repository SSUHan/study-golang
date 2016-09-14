package main 

import "fmt"
import "time"
/* 
	Golang 은 여러 채널을 손쉽게 이용할 수 있도록 select 분기문을 제공한다
	각 채널에 값이 들어오면 해당 case 가 실행된다 : close 함수로 채널을 닫았을때도 case 는 실행된다
	보통 select 를 for 를 이용해 반복해준다
	default 에 적절한 처리를 하지 않으면, CPU 코어를 모두 잡아먹으므로 주의한다
*/

func main(){
	c1 := make(chan int)
	c2 := make(chan string)

	go func(){
		for{
			c1<-10 // 계속 값을 집어넣는 고루틴
			fmt.Println(">> input c1")
			time.Sleep(100*time.Millisecond)
		}
	}()

	go func(){
		for{
			c2<-"Hello MMM" // 마찬가지
			fmt.Println(">> input c2")
			time.Sleep(500*time.Millisecond)
		}
	}()

	go func(){
		for{
			select{
			case i:= <-c1:
				// c1 에 값이 들어오면
				fmt.Println("c1 : ", i)
			case i := <-c2:
				fmt.Println("c2 : ", i)
			case <-time.After(50*time.Millisecond):
				fmt.Println("timeout")
			}
		}
	}()

	// select 분기문은 채널에 값을 보낼 수 도 있다
	// select 에서 채널에 값을 보내는 경우, 계속 값을 넣다가, 다른 case가 해당될시 그것을 실행한다


	fmt.Scanln()  // 엔터누르면 종료

}