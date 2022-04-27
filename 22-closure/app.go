package main

import "fmt"

func main() {
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	nextInt2 := intSeq()
	fmt.Println(nextInt2())
	fmt.Println(nextInt2())
	fmt.Println(nextInt2())

	seqString := myinitSeq()
	fmt.Println(seqString())
	fmt.Println(seqString())
	fmt.Println(seqString())

	fmt.Println(myinitSeq()())
	fmt.Println(myinitSeq()())

}

func intSeq() func() int {
	i:=0
	return func() int {
		i++
		return i
	}
}

func myinitSeq() func() string {
	y:=0
	return func() string {
		y++
		return fmt.Sprintf("y = %d",y)
	}
}