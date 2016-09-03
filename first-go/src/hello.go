package main

import "fmt"


var a = "hello2"

func main(){
	var x,y int = 30, 50
	var age, name = 40, "Brian Lee"
	test := "hello"
	
	_ = age // 사용하지 않는 변수로 인해서 생기는 컴파일을 방지
	_ = test // 빠른 선언도 이렇게 해서 할 수 있다.
	
	fmt.Println(name)
	fmt.Println(x, y, name)
}