package main

import (
	"fmt"
)

func main() {
	fnfor();
	fnWhile();
}

func fnfor(){
	for i:=0;i<=10;i++ {
		// fmt.Println(i);
		fmt.Printf("index > %d\n",i);
	}
}
func fnWhile(){
	index := 0;
	for index < 10 {
		fmt.Printf("index >> %d\n",index);
		index++;
	}
}