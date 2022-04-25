package main

import "fmt"

func main() {
	fmt.Println("get started app")
	defer fmt.Println("processing...")
	fmt.Println("end off app")

	for i:=0; i<10; i++ {
		fmt.Println(">",i)
	}

	for i:=0; i<10; i++ {
		defer fmt.Println(">",i)
	}
}