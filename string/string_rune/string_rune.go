package main

import "fmt"

func main() {
	var char rune = '한' // rune 타입으로 한문자 담기

	fmt.Printf("%T\n", char) // char 타입 출력
	fmt.Println(char)        // char 값 출력
	fmt.Printf("%c\n", char) // 문자 출력
}
