package main

import "fmt"
import "strconv"

func formatString(arg interface{}) string {
	switch arg.(type){
	case int:
		i := arg.(int)
		return strconv.Itoa(i)
	case float32:
		f := arg.(float32)
		return strconv.FormatFloat(float64(f), 'f', -1, 32)
	case float64:
		f := arg.(float64)
		return strconv.FormatFloat(f, 'f', -1, 64)
	case string:
		s := arg.(string)
		return s
	default :
		return "unknown type"
	}
}

func main(){
	fmt.Println(formatString(3))
	fmt.Println(formatString(2.6))
	fmt.Println(formatString("hellowlrod"))
}