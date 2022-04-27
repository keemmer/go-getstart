package main

import (
	"fmt"
	"time"
)

func main() {

	// run1()
	// run2()

	go run1()
	go run2()
	
	time.Sleep(5 * time.Second)
	
}


func run1(){
	for i := 0; i < 100; i++ {
		fmt.Println("run1 loop")
	}
}

func run2(){
	for i := 0; i < 100; i++ {
		fmt.Println("run2 loop")
	}
}