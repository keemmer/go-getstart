package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("hello route")
	fmt.Fprintf(w, "hello world keemmer\n")
}

func main() {
	http.HandleFunc("/hello", hello)

	port := 8080
	fmt.Printf("server started on %v", port)
	http.ListenAndServe(fmt.Sprint(":", port), nil)
}
