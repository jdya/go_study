package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		stdin.ReadString('\n')
	}
	return n, err

}

func main() {
	rand.Seed(time.Now().UnixNano())
	
	r := rand.Intn(100)
	cnt := 1
	for {
		fmt.Printf("input number: ")
		n, err := InputIntValue()
		if err != nil {
			fmt.Println("Only Number! ")
		} else {
			if n > r {
				fmt.Println("입력한 숫자가 더 크다")
			} else if n < r {
				fmt.Println("입력한 숫자가 더 작다")

			} else {
				fmt.Println("good job!. you tried:", cnt)
				break
			}
			cnt++
		}
	}
}
