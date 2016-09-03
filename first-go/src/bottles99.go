package main

import (
	"fmt"
	"strconv"
)
// fmt.Printf()
// fmt.Println() 에서도 서식문자 사용이 가능하다
func main(){
	for i := 99; i>=1;{
		if i == 1{
			fmt.Println("%d bottles of beer on the wall, "+strconv.Itoa(i)+" bottle of beer", i)
		}else{
			fmt.Println(strconv.Itoa(i)+" bottles of beer on the wall, "+strconv.Itoa(i)+" bottles of beer")
		}
		if i>2{
			i--
			fmt.Println("Take one down, pass it around, "+strconv.Itoa(i)+" bottles of beer on the wall")
		}else if i == 2{
			i--
			fmt.Println("Take one down, pass it around, "+strconv.Itoa(i)+" bottle of beer on the wall")	
		}else if i == 1{
			i--
			fmt.Println("Take one down, pass it around, No more bottles of beer on the wall")
			fmt.Println("No more bottles of beer on the wall, No more bottles of beer")
			fmt.Println("Go to the store and buy some more, 99 bottles of beer on the wall")	
		}

	} 
}