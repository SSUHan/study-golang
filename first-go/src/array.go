package main

import (
	"fmt"
)

func main(){
	a :=[5] int {32, 3,12, 34,22}
	fmt.Println(a) // 그냥 뽑히네

	for i, value := range a{ // 첫번째 리턴값으로 인덱스 정보를 무조건 받아야 한다
		fmt.Println(i ,value)
	}

	for _ , value := range a{ // 인덱스 정보를 사용하지 않을 것이라면 밑줄로 받는다
		fmt.Println(value)
	}
}