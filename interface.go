
package main

import "fmt"

type SalaryCalculator interface {
	calculateSalary() int
}

type Permanent struct {
	empId    int
	basicPay int
	pf       int
}

type Contract struct {
	empId    int
	basicPay int
}

func (p Permanent) calculateSalary() int {
	return p.basicPay + p.pf
}

func (c Contract) calculateSalary() int {
	return c.basicPay
}

func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense += expense + v.calculateSalary()
	}
	fmt.Println("Total expense is", expense)
}

func main() {
	pemp1 := Permanent{1, 5000, 20}
	pemp2 := Permanent{2, 6000, 30}
	cemp1 := Contract{3, 3000}
	employees := []SalaryCalculator{pemp1, pemp2, cemp1}
	fmt.Println(employees)
	totalExpense(employees)

}
