package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/bar", barHandler) // "/bar" 핸들러 등록
	http.ListenAndServe(":3000", nil)
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()    // 쿼리 인수 가져오기
	name := values.Get("name") // 특정 키 값이 있는지 확인
	if name == "" {
		name = "World"
	}
	id, _ := strconv.Atoi(values.Get("id")) // id 값을 가져와서 int 타입변환
	fmt.Println("Server is started!")
	fmt.Fprintf(w, "Hello %s! id:%d", name, id)

}
