package main

import "fmt"

func main() {
	var u user
	u.fname = "keemmer"
	u.lname = "hello"
	u.age = 30
	fmt.Printf("struct > %v\n", u.fname)
	show(u)
	update(&u)
	show(u)

	fullname := u.getFullname()
	fmt.Printf("struct > %v\n", fullname)
	u.showFullname();

	u = u.decreaseAge(4)
	show(u)


}

type user struct {
	fname string
	lname string
	age   int
}

func (u user) getFullname() string {
	fmt.Printf("%v %v\n", u.fname, u.lname)
	return u.fname + " " + u.lname
}
func (u user) showFullname() {
	fmt.Printf("show fullname > %v %v\n", u.fname, u.lname)
}
func (u user) decreaseAge(num int) user {
	u.age = u.age - num
	return u
}
func show(u user) {
	fmt.Printf("struct > %v\n", u)
}

func update(u *user) {
	u.age = 31
	u.fname = "keemmerTH"
	fmt.Printf("struct pointer > %v\n", u)
	fmt.Printf("struct pointer > %v\n", *u)
}
