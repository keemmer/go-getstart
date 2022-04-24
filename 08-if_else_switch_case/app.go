package main

import "fmt"

func main() {
	value := 10
	value2 := 20

	if value == 10 {
		fmt.Printf("data = %d\n", value)
	} else {
		fmt.Printf(">>data = %d\n", value)
	}

	if value == 10 && value2 == 20 {
		fmt.Printf("value = %d && value2 = %d\n", value, value2)
	} else {
		fmt.Printf(">>value = %d && value2 = %d\n", value, value2)
	}

	if result := doSomething(); result == "ok" {
		fmt.Println("ok")
	} else {
		fmt.Println("not ok")
	}

	fnSwitchCase(2)
}

func doSomething() string {
	return "ok"
}

func fnSwitchCase(index int) {
	switch index {
	case 0:
		fmt.Println("case 0")
		break
	case 1:
		fmt.Println("case 1")
		break
	case 2:
		fmt.Println("case 2")
		break
	default:
		fmt.Println("some thing case")
	}
}
