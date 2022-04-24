package main

import "fmt"

func main() {
	numbers := 100
	temp := &numbers



	fmt.Printf("nuber %v\n", numbers)
	fmt.Printf("temp %v\n", temp)
	fmt.Printf("temp %v\n", *temp)


	*temp = 200
	
	fmt.Printf("nuber %v\n", &numbers)
	fmt.Printf("nuber %v\n", numbers)
	fmt.Printf("temp %v\n", temp)
	fmt.Printf("temp %v\n", *temp)


	var msg string = "keemmer pointer"
	var msgPointer* string = &msg

	fmt.Println(&msg);
	fmt.Println(msg);
	fmt.Println(msgPointer);
	fmt.Println(*msgPointer);

	*msgPointer = "data keemer"

	fmt.Println(&msg);
	fmt.Println(msg);
	fmt.Println(msgPointer);
	fmt.Println(*msgPointer);

	msg = "msg set data!"

	fmt.Println(&msg);
	fmt.Println(msg);
	fmt.Println(msgPointer);
	fmt.Println(*msgPointer);

}