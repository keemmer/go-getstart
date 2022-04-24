package main

import "fmt"

func main() {
	var array1 []int = []int{1, 2, 3, 4}
	var array2 = []int{2, 3, 4, 5, 6}
	var array3 [3]string
	fmt.Println(array1[0])

	for _, item := range array1 {
		fmt.Printf("%d \n", item)
	}
	for idx, item := range array2 {
		fmt.Printf("%d %v \n", item, idx)
	}

	for _, item := range array3 {
		fmt.Printf("- %v \n", item)
	}

	array3[0] = "golang"
	array3[1] = "nodejs"
	array3[2] = "python"

	for _, item := range array3 {
		fmt.Printf("- %v \n", item)
	}
}
