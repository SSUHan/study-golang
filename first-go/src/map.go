package main

import (
	"fmt"
)

func main(){
	// map 을 할당하는 5가지 방법
	var a map[string]int = make(map[string]int)

	var b = make(map[string]int)

	c := make(map[string]int)

	d := map[string]int{"hello":10, "no":5}

	e := map[string]int{
		"hello":20,
		"no":10, // 여러 줄로 나열할 때는 마지막 요소에 콤마르 써야한다 (특이)
	}

	_ = a
	_ = b
	_ = c
	_ = d

	fmt.Println(e["hello"])

	// 맵과 조건문
	if value, ok := e["hello2"]; ok{ // map 의 두번째 리턴값은 값의 존재 여부이다. 이를 조건문 한줄로 처리하는 방법이다
		fmt.Println(value)
	}else{
		fmt.Println("no value")
	}

	// 맵과 순환
	for key, value := range e{ // range 로 뽑으면 키, 값 이 리턴, 사용하고 싶지 않다면 _ 쓴다
		fmt.Println(key, value)
	}

	// 맵과 삭제
	fmt.Println(e)
	delete(e, "hello")
	fmt.Println(e)

	// 맵안에 맵 을 만들 수 있다
	// := map[key]map[key]value
}