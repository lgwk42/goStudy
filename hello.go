package main

import "fmt"

func main() {

	//변수 선언 var 변수명 타입 순으로 선언
	//사용하지 않을 경우 에러 발생
	var a int = 10
	var msg string = "Hello"

	//변수 값 변경
	a = 20
	msg = "Good Morning"

	//출력문
	fmt.Println(msg, a)
}
