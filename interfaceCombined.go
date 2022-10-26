
package main

import "fmt"

type SalaryCalculator interface {
	CalculateSalary() int
}

type Leaves interface {
	CalculateLeavesLeft() int
}

type EmployeeOperations interface {
	SalaryCalculator
	Leaves
}

type Employee struct {
	salary      int
	pf          int
	totalLeaves int
	leftLeaves  int
}

func (e Employee) CalculateSalary() int {
	return e.salary + e.pf
}

func (e Employee) CalculateLeavesLeft() int {
	return e.totalLeaves - e.leftLeaves
}

func main() {
	e1 := Employee{5000, 50, 10, 7}
	var op Employee = e1
	fmt.Println(op.CalculateSalary())
	fmt.Println(op.CalculateLeavesLeft())
}
