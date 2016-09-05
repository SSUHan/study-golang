package main 

import "fmt"

type Rectangle struct{
	height int
	width int
}

// 구조체를 받는 함수의 변수 부분을 리시버라고 한다
// 리시버가 구조체의 포인터를 받으면, 그 자체 값을 바꿧을때 영향을 받는다
func (rect * Rectangle) scaleA(factor int){
	rect.width = rect.width * factor
	rect.height = rect.height * factor
}

// 그러나 리시버가 구조체의 인스턴스만 받는다면, 값을 복사해서 오기 때문에 기존값을 변화시키지 않는다
func (rect Rectangle) scaleB(factor int){
	rect.width = rect.width * factor
	rect.height = rect.height * factor
}
func main(){
	rect1 := Rectangle{30, 50}
	rect1.scaleA(10)
	fmt.Println(rect1)

	rect2 := Rectangle{30, 50}
	rect2.scaleB(10)
	fmt.Println(rect2)
}