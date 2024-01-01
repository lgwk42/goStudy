//go run "파일명"
//go mod init "모듈명" +모듈명은 파일 경로 일치하게 하는것이 좋음
//go build -> ./파일명 = 빌드 파일 생성, 실행

package main

import "fmt"

func main() {

	//변수 선언 var 변수명 타입 순으로 선언
	//사용하지 않을 경우 에러 발생
	var a int = 10           //정수
	var msg string = "Hello" //문자열

	//변수 값 변경
	a = 20
	msg = "Good Morning"

	//출력문
	fmt.Println(msg, a)
}
