package main

import "fmt"

func main() {
	number := []int{1, 2, 3,4,5,6,7,8,9,10}

	// leading remove
	showSlice(number)
	number = number[1:len(number)]
	showSlice(number)
	number = number[1:len(number)]
	showSlice(number)
	// tailing remove
	number = number[0:len(number)-1]
	showSlice(number)
	number = number[0:len(number)-1]
	showSlice(number)
	
	//remove specific removeIndex
	number = removeIndex(number,3)
	showSlice(number)
	// fmt.Printf(" >> %v\n",number[3:])

}
func showSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func removeIndex(s []int,index int) []int{
	// return append(s[:index],s[index+1:][0])
	return append(s[:index],s[index+1:]...)
}