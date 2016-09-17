// 계산 패키지
package calc

// 전역 변수도 됨
var TotalNum int

// 두 정수를 더함
// 두 줄도 됨
// 대문자로 시작하면 public : 외부에서 사용할 수 있음
func Sum(a int, b int) int {
	return a + b
}

// 두 정수를 뺌 
// 소문자로 시작하면 private : 파일 내부에서만 사용이 가능
func sub(a int, b int) int {
	return a - b
}
