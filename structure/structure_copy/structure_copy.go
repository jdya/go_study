package main

import "fmt"

type Student struct { // 구조체 선언
	Age   int
	No    int
	Score float64
}

func PrintStudent(s Student) {
	fmt.Printf("나이: %d 번호: %d 점수: %.2f\n", s.Age, s.No, s.Score)
}

func main() {
	var student = Student{15, 23, 88.2} // 변수 초기화

	student2 := student // 구조체 복사

	PrintStudent(student2) // 함수 호출시에도 구조체 복사
}
