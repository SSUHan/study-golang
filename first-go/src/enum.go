package main

import(
	"fmt"
)

func main(){
	// const(
	// 	Sunday = 0
	// 	Monday = 1
	// 	Tuesday = 2

	// )

	const (
		Sunday = iota // 0 부터 시작 만약 1부터 시작하고 싶다면 iota + 1 이렇게 쓰면 된다
		Monday 
		Tuesday
		Wednesday
		Thursday
		Friday
	)

	fmt.Println(Friday)
	
	const (
		a = 1 << iota // a == 1
		b = 1 << iota // b == 2
		c = 1 << iota // c == 4
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}