package main

import (
	"fmt"
	"unsafe"
)

// 메모리 낭비를 줄이려면
// 8바이트보다 작은 필드는 8바이트 크기를 고려해서 몰아서 배치한다
type User struct {
	A int8 // 1바이트
	C int8 // 1바이트
	E int8 // 1바이트
	B int  // 8바이트
	D int  // 8바이트
}

func main() {
	user := User{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(user))
}

// 결합도 coupling : 모듈간 상호 의존관계 형성해서 서로 강하게 결합되어 있는 정도
// 응집도 cohesion : 모듈의 완성도. 모듈 내부의 모든 기능이 단일 목적에 충실하게 모여있는가
