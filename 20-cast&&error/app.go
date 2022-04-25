package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var index int8 = 15
	var bigIndex int32
	bigIndex = int32(index)
	fmt.Printf("type: %T\n", index)
	fmt.Printf("type: %T\n", bigIndex)

	a := strconv.Itoa(108)
	fmt.Printf("%T\n", a)


	lines_yesterday := "50"
	lines_today := "108o"

	yesterday, err := strconv.Atoi(lines_yesterday)
	if err != nil {
		log.Fatal(err)
	}

	today, err := strconv.Atoi(lines_today)
	if err != nil {
		log.Fatal(err)
	}
	lines_more := today - yesterday

	fmt.Println(lines_more)
	
	fmt.Printf("type: %T\n", lines_today)
	fmt.Printf("type: %T\n", yesterday)
}
