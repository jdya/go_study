package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){ // 함수 리터럴: 함수명을 적지 않고 함수 타입 변숫값으로 대입되는 함숫값을 의미함. 함수명이 없기 때문에 함수명으로 직접 함수를 호출할 수 없고 함수 타입 변수로만 호출됨.
		fmt.Fprint(w, "Hello World")
	})
}
http.ListenAndServer(":3000",nil)