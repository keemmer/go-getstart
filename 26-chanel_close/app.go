package main

import "fmt"

func routine1(ch chan int, countTo int) {
	for i := 0; i < countTo; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int, 1)
	go routine1(ch, 10)
	// for true {
	// 	value, ok := <-ch
	// 	if !ok {
	// 		fmt.Println("No more data")
	// 		break
	// 	}
	// 	fmt.Println(value)
	// }
	for value := range ch {
		fmt.Println(value)
	}
	fmt.Println("No more data")
}