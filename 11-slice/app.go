package main

import "fmt"

func main() {
	number := make([]int,3,5)
	number = append(number,1)
	number = append(number,2)
	number = append(number,3)

	showSlice(number)

	// var number2 []int
	// var number2 []int = []int{1,2,3,4}
	var number2 []int
	number2 = append(number2,1)
	number2 = append(number2,2)
	number2 = append(number2,3)
	number2 = append(number2,4)
	showSlice(number2)
}

func showSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x),cap(x),x)
}