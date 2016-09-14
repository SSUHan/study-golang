package main

import (
	"fmt"
	"github.com/golang/example/stringutil" // github에 있는 url 을 쓴다
	"calc"
)

func main(){
	fmt.Println(stringutil.Reverse("olleh"))
	fmt.Println(calc.Sum(1,2))
}