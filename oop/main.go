package main

import (
	"oop/employee"
)

func main() {
	e := employee.New("Prahar", 10, 3)
	e.LeavesRemaining()
}
