package main

import "fmt"

func main() {
	company := 0
	// Func 1
	func() {
		fmt.Println("Anon Function")
	}()
	
	// Func 2
	counter := func() int {
		comapny += 1
		return company
	}
	
	// Func 3
	func(name string) {
		fmt.Println(name)
	}("Telecom")

	fmt.Println(counter())
	fmt.Println(comapny)
}
