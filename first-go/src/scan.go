package main

import "fmt"

func main() {
	var s1, s2 string
	n, _ := fmt.Scan(&s1, &s2) // 공백과 앤터키로 구분하여 입력받는다
	fmt.Println("입력 갯수 : ", n) // 두번째 인자로는 nil 이 들어오는데 정체는 아직 모르겠다
	fmt.Println(s1, s2)

	n, _= fmt.Scanln(&s1, &s2) // 공백으로만 구분하여 입력받는다, 엔터키를 누르면 입력이 끝난다
	fmt.Println("입력 갯수 : ", n)
	fmt.Println(s1, s2)

	var num1, num2 int

	n, _ = fmt.Scanf("%d,%d", &num1, &num2) // 형식을 지정해서 입력을 받을 수도 있다
	// %d,%d 로 입력을 받았기 때문에, 입력받을때 콤마로 구분하여 입력을 받는다
	fmt.Println("입력 갯수 : ", n)
	fmt.Println(num1, num2)

}