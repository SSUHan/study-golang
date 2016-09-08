package main

import "fmt"

// 각 값이나 인스턴스의 실제 타입은 상관하지 않고, 
// 구현된 메서드로만 타입을 판단하는 방식을 "덕 타이핑" 방식이라고 한다

type Duck struct{

}
func (d Duck) speak(){
	fmt.Println("꽥!!!!")
}
func (d Duck) feature(){
	fmt.Println("오리는 오리지 뭐")
}


type Person struct {

}
func (p Person) speak(){
	fmt.Println("나는 사람인데 오리처럼 어떻게 움")
}
func (p Person) feature(){
	fmt.Println("나는 손이 두개 발이 두개인 사람임")
}

type whoer interface{
	speak()
	feature()
}

func whoareU(who whoer){ // 여기에 누가 오든, 해당 인터페이스의 메서드를 가지고만 있다면 부를 수 있다
	who.speak()
	who.feature()
}

func main(){
	var duck Duck
	var p1 Person

	whoareU(duck)
	whoareU(p1)
}