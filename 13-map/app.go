package main

import "fmt"

func main() {
	var color = map[string]int{"red": 0, "green": 1, "blue":3}
	fmt.Printf("data: %v\n", color)
	fmt.Printf("data: %v\n", color["red"])

	var data = make(map[string]string)
	data["web"] = "reactjs"
	data["embeded"] = "python"
	data["backend"] = "golang"
	fmt.Printf("%v\n", data)
	fmt.Printf("%v\n", data["backend"])


	var languages = make(map[string]map[string]int)
	languages["python"] = map[string]int{"price": 200}

	languages["nodejs"] = make(map[string]int)
	languages["nodejs"]["price"]=300
	languages["nodejs"]["code"]=444
	
	fmt.Printf("%v\n", languages)
	fmt.Printf("%v\n", languages["nodejs"]["code"])
}