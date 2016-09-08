package main

import "fmt"

type hello interface{
	// empty
}
// 

type myInt int // int 형을 myInt 로 정의

func (i myInt) myPrint(){ // myInt 에 myPrint 메서드를 연결
	fmt.Println("my", i)
}

type Rectangle struct{ // Rectangle 구조체를 정의하고
	width, height int
}

func (r Rectangle) myPrint(){ // 같은 이름으로 사각형 구조체에 연결
	fmt.Println(r.width, r.height)
}

type Printer interface{ // myPrint를 부르는 인터페이스를 정의
	myPrint()
}

func main(){
	var h hello
	fmt.Println(h)

	var i myInt = 5
	r := Rectangle{20,30}
	var p Printer // 인터페이스 선언

	p = i // i를 인터페이스 p 에 대입
	p.myPrint()

	p = r // r를 인터페이스 p 에 대입
	p.myPrint()

	// 선언하면서 초기화 할 수 도 있다
	p2 := Printer(i)
	p2.myPrint()
	p2 = r // 이렇게 대입 가능
	p2.myPrint()

	// 배열(슬라이스) 형태로도 인터페이스 선언 가능
	pArr := []Printer{i, r}

	for index, value := range pArr{
		fmt.Println(">> ", value)
		pArr[index].myPrint()
	}

	for _ , value := range pArr{
		value.myPrint()
	}
}