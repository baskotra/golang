package main

import "fmt"

func squares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		n := number % 10
		number = number / 10
		sum += n * n
	}
	squareop <- sum
}

func cubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		n := number % 10
		sum += n * n
		number = number / 10
	}
	cubeop <- sum
}

func main() {
	number := 595
	sqrch := make(chan int)
	cubech := make(chan int)

	go squares(number, sqrch)
	go cubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Sum is : ", squares+cubes)
}
