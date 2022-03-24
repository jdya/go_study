package main

import "fmt"

func main() {
	var slice = []int{1, 2, 3}
	slice2 := append(slice, 4)
	var slice3 []int

	fmt.Println(slice)
	fmt.Println(slice2)

	for i := 1; i <= 10; i++ {
		slice3 = append(slice3, i)
	}

	slice3 = append(slice3, 11, 12, 13, 14, 15)
	fmt.Println(slice3)
}
