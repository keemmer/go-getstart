package main

import (
	"demo/employee"
)

func main() {
	// e := employee.Employee{
	// 	FirstName:   "Sam",
	// 	LastName:    "Adolf",
	// 	TotalLeaves: 30,
	// 	LeavesTaken: 20}
	e := employee.Init("Sam", "Adolf", 30, 20)
	e.LeavesRemaining()
	e = employee.Init("keem", "Adolf", 30, 20)
	e.LeavesRemaining()
}
