package main 

import "fmt"
import "runtime"
import _"time"
// 버퍼 채널을 만드는 방법은 make(chan <type>, <buffer size>)

func main(){
	runtime.GOMAXPROCS(1)

	done := make(chan bool, 2) // 버퍼 크기가 2 인 타입 bool 의 채널을 생성
	count := 4

	go func(){
		for i:=0; i<count; i++{
			done <- true
			fmt.Println("고루틴 : ", i)
			// time.Sleep(1*time.Second)
		}
	}()

	for i:=0; i<count; i++{
		<- done
		fmt.Println("메인 함수 : ",i)
	}

	// range 와 close 를 사용하는 것
	c := make(chan int)
	go func(){
		for i:=0; i<5; i++{
			c<-i
		}
		close(c)
	}()
	for i:= range c{
		fmt.Println("close 메인함수 : ", i)
	}
	// 즉 range chan 은 채널에 값이 몇개나 들어올지 모를때, 값이 들어올때마다 꺼내기 위해서 사용한다
	
	// 채널에서 값을 가져온 것의 두번째 리턴값은 채널이 닫혔는지, 열렸는지에 대한 정보을 얻을 수 있다
	// true : open channel, false : close channel
	c2 := make(chan int)
	go func(){
		c2<-10
	}()

	a, is_open := <-c2
	fmt.Println(a, is_open)

	close(c2)
	a, is_open = <-c2
	fmt.Println(a, is_open)
}