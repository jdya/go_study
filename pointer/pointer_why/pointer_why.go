// 변수 대입, 함수 전달 - 값 복사 --> 메모리 공간 많이 차지 --> 복사할때 성능문제
// 그리고 다른 공간으로 복사 --> 변경사항 적용 안됨

package main

import "fmt"

type Data struct {
	value int
	data  [200]int
}

func ChangeData(arg *Data) { // 매개변수로 Data 포인터를 받음 (메모리 주소값인 8바이트만 복사됨)
	// arg 포인터 변수가 data 변수의 메모리 주소를 값으로 가지고 있어서 Data 구조체의 내부 필드값 변경 가능
	arg.value = 999     // arg 데이터를 변경
	arg.data[100] = 999 // arg 데이터를 변경
}

func main() {
	var data Data

	ChangeData(&data)
	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100])

}
