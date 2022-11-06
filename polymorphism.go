package main

import "fmt"

type Income interface {
	calculate() int
	source() string
}

type fixedBilling struct {
	projectName  string
	biddedAmount int
}

type timeAndMaterial struct {
	projectName string
	noOfHours   int
	hourlyRate  int
}

func (fb fixedBilling) calculate() int {
	return fb.biddedAmount
}

func (fb fixedBilling) source() string {
	return fb.projectName
}

func (tm timeAndMaterial) calculate() int {
	return tm.hourlyRate * tm.noOfHours
}

func (tm timeAndMaterial) source() string {
	return tm.projectName
}

func calculateNetIncome(ic []Income) {
	var netIncome int = 0
	for _, income := range ic {
		fmt.Printf("Income from source %s = ₹%d\n", income.source(), income.calculate())
		netIncome += income.calculate()
	}
	fmt.Println("Net Income form firm is = ", netIncome)
}

func main() {
	project1 := fixedBilling{
		"Bharamala Project",
		2000,
	}
	project2 := timeAndMaterial{
		"Wall Building",
		10,
		205,
	}

	incomeStreams := []Income{project1, project2}
	calculateNetIncome(incomeStreams)
}
