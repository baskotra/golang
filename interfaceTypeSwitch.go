package main

import "fmt"

type Describer interface {
	Describe()
}

type Person struct {
	name string
	age  int
}

func (p Person) Describe() {
	fmt.Printf("Person name is %v and age is %v", p.name, p.age)
}

func findType(i interface{}) {
	switch v := i.(type) {
	case Describer:
		fmt.Printf("Type is %T\n", v)
		v.Describe()
	default:
		fmt.Println("Not suitable interface")
	}
}

func main() {
	findType("Prakhar")
	pra := Person{"Prakhar", 23}
	findType(pra)

}
