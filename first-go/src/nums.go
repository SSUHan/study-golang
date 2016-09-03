package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main(){
	// 1. 실수 선언 / 비교
	var a float64 = 10.0

	for i:=0; i<10; i++{
		a = a - 0.1
	}

	fmt.Println(a)
	fmt.Println(a == 9.0) // 이렇게 비교했을때, true 가 나와야 할 것 같지만, 실제로는 false 가 나온다

	// 따라서 앱실론을 써서
	const epsilon = 1e-14
	fmt.Println(math.Abs(a-9.0) <= epsilon)

	// 2. 복소수
	var num1 complex64 = 1+2i
	var num2 = 2+3i

	fmt.Println(num1, num2)

	var r1 = real(num1) // 실수부 찾기
	var i1 = imag(num1) // 허수부 찾기

	fmt.Println(r1, i1)

	// 3. 바이트 (byte)

	var char byte = 'a'
	fmt.Println(char)

	// 4. 오버플로우, 언더플로우(OverFlow, UnderFlow)
	var max8u uint8 = math.MaxUint8
	// var max8u uint8 = math.MaxUint8 + 1  // overflow 
	var max16u uint16 = math.MaxUint16
	var max32u uint32 = math.MaxUint32
	var max64u uint64 = math.MaxUint64
	fmt.Println(max8u)
	fmt.Println(max16u)
	fmt.Println(max32u)
	fmt.Println(max64u)

	 //var underflow uint8 = 0 -1 // underflow

	// 5. 변수의 크기 구하기
	fmt.Println(unsafe.Sizeof(max8u))
	fmt.Println(unsafe.Sizeof(max16u))
	fmt.Println(unsafe.Sizeof(max32u))
	fmt.Println(unsafe.Sizeof(max64u))

}