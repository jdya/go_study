package main

import (
	"fmt"
)

func main() {
	str1 := "Hello 월드!" // 한영문자가 섞인 문자열
	arr := []rune(str1) // []rune 으로 타입 변환

	str2 := "Hello 고랑!" // 한영문자가 섞인 문자열

	for i := 0; i < len(arr); i++ { // 문자열 크기를 얻어 순회
		fmt.Printf("타입: %T 값: %d 문자값: %c\n", arr[i], arr[i], arr[i])
	}

	for _, v := range str2 { // range를 이용한 순회
		fmt.Printf("타입: %T 값: %d 문자값: %c\n", v, v, v)
	}
}
