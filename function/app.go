package main

import "fmt"

func main() {
	fn1("GO LANG")
	display()
	fn3("javascript", 12)
	const a, b = 10, 20
	fmt.Printf("%d + %d = %d\n", a, b, sum(a, b))

	// var x,y int = swap(a,b);
	x, y := swap(a, b)
	fmt.Printf("a=%d,b=%d <> a=%d,b=%d\n", a, b, x, y)

	x, y = swapV2(23, 34)
	fmt.Printf("a=%d,b=%d <> a=%d,b=%d\n", 23, 34, x, y)
	_x, _y, title := swapV3(23, 34)
	fmt.Printf("a=%d,b=%d <> a=%d,b=%d,title=%s \n", 23, 34, _x, _y, title)

}

func display() {
	fmt.Println("fn: keemmer")
}

func fn1(title string) {
	fmt.Println(title)
}

func fn3(lang string, version int) {
	fmt.Printf("lang : %s = version: %d\n", lang, version)
}

func sum(a int, b int) int {
	return a + b
}

func swap(a, b int) (int, int) {
	return b, a
}

func swapV2(a, b int) (x, y int) {
	y = a
	x = b
	return
}

func swapV3(a, b int) (x, y int, title string) {
	y = a
	x = b
	title = "v3"
	return
}
