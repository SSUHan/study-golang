package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

type Data struct{ // Data 구조체 정의
	tag string    // 풀의 태그 역할
	buffer []int  // 내용물
}

func main(){
	runtime.GOMAXPROCS(runtime.NumCPU())

	pool := sync.Pool{
		New: func() interface{} {	// Get 함수를 사용했을때 호출될 함수를 정의
			data := new(Data)		// 새로 메모리 할당
			data.tag = "new"		// 새로운 풀에는 tag = new
			data.buffer = make([]int, 10) // 슬라이스에 공간 할당
			fmt.Println(">> int Pool New..")
			return data
		}, // 여기 콤마 있어야 합니다. 왜지?
	}

	for i:=0; i< 10; i++{
		go func(a int) { // GoRoutine 10개 생성
			data := pool.Get().(*Data) // pool 에서 *Data 타입으로 데이터를 가져옴. Go에서는 캐스팅을 저런식으로 하나봄

			for index := range data.buffer{
				data.buffer[index] = rand.Intn(100)
			} // data의 buffer 에 랜덤값 채워넣기
			fmt.Println("input random : ", data, a)
			data.tag = "used" // 사용되었다
			pool.Put(data) // 다시 넣는다
		}(i)
	}

	for i:= 0; i<10;i++{
		go func(a int){
			data := pool.Get().(*Data)

			n := 0
			for index := range data.buffer{
				data.buffer[index] = n
				n +=2
			}
			fmt.Println("change 2 : ", data, a)
			data.tag = "used"
			pool.Put(data)
		}(i)
	}
	fmt.Scanln()
}

