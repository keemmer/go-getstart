package main

import (
	"fmt"
)

func main() {
	fnfor()
	fnWhile()
	fnWhileUsingBreak()
}

func fnfor() {
	for i := 0; i <= 10; i++ {
		// fmt.Println(i);
		fmt.Printf("index > %d\n", i)
	}
}
func fnWhile() {
	index := 0
	for index < 10 {
		fmt.Printf("index >> %d\n", index)
		index++
	}
}

func fnWhileUsingBreak() {
	index := 0
	for true {
		if index > 3 {
			break
		}
		fmt.Printf("while break : %d\n", index)
		index++
	}
}
