package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int,5)
	go routine1(ch,1)
	go routine1(ch,2)
	go routine1(ch,3)


	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	time.Sleep(1*time.Second)
}

func routine1(ch chan int,payload int) {
	ch <- payload
	// fmt.Println("step 1")
	// ch <- payload
	// fmt.Println("step 2")
	// ch <- payload
	// fmt.Println("step 3")
}
