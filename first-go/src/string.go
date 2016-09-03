package main

import(
	"fmt"
	"unicode/utf8"
)

func main(){
	var s1 string = "hello world"
	
	// 여러줄의 경우에는 `(백쿼트) 를 써서 표현합니다
	var s2 string = `안녕하세요
건프 입니다.
	껄껄껄`

	var s3 = "안녕하세요"
	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Println(len(s1))
	fmt.Println(len(s3))
	fmt.Println(utf8.RuneCountInString(s3)) // 유니코드 문자열의 실제 길이를 알고 싶을때는 이렇게 해야한다
}