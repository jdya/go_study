package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux() // ServeMux 인스턴스 생성

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World") // 인스턴스에 핸들러 등록
	})
	mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Bar")
	})

	http.ListenAndServe(":3000", mux) // mux 인스턴스 사용

	// mux
	// multiplexer의 약자
	// 웹서버는 각 url에 해당하는 핸들러들을 등록한 다음 http 요청이 왔을때 url에 해당하는 핸들러를 선택해서 실행하는 방식임
	// 비슷한 의미: 라우터 (router)
}
