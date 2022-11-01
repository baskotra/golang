package employee

import "fmt"

type employee struct {
	firstName   string
	totalLeaves int
	leavesTaken int
}

func (e employee) LeavesRemaining() {
	fmt.Printf("Total leaves pending for %s are: %d", e.firstName, e.totalLeaves-e.leavesTaken)
}

func New(firstName string, totalleaves int, leavestaken int) employee {
	e := employee{firstName, totalleaves, leavestaken}
	return e
}
