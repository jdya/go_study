package main

import (
	"container/list"
	"fmt"
)

type Queue struct { // 리스트를 이용한 큐 구조체 정의. 내부 필드로 리스트 갖고 있음.
	v *list.List
}

func (q *Queue) Push(val interface{}) { // Push() 메서드는 리스트의 PushBack() 메서드를 이용해서 맨 뒤에 요소를 추가한다. 빈 인터페이스를 이용하여 모든 타입의 데이터를 저장함
	q.v.PushBack(val)
}

func (q *Queue) Pop() interface{} { // 요소를 반환하면서 삭제
	front := q.v.Front()
	if front != nil {
		return q.v.Remove(front)
	}
	return nil
}

func NewQueue() *Queue {
	return &Queue{list.New()}
}

func main() {
	queue := NewQueue() // 새로운 큐 생성

	for i := 1; i < 5; i++ { // 요소 입력
		queue.Push(i)
	}
	v := queue.Pop()
	for v != nil { // 요소 출력
		fmt.Printf("%v -> ", v)
		v = queue.Pop()
	}
}
