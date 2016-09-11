package main 

import "fmt" 

// send-only channel
// receive-only channel 

func producer(c chan<- int){ // send-only channel 
	for i:=0; i<5; i++{
		fmt.Println("producer : ", i)
		c<-i
	}
	c<-100
	// <-c // 꺼내려 들면 이러면 컴파일 오류
}

func consumer(c <-chan int){ // receive-only channel
	for i := range c{ // 꺼내기만 합니다 
		fmt.Println("consumer : ", i)
	}
	// c<- 100 // 넣으려 들면 컴파일 오류
}
func main(){	
	c := make(chan int)

	go producer(c)
	go consumer(c) // 인자로 넘길때는 특별한 행위를 하지 않아도 된다

	fmt.Scanln()
}

