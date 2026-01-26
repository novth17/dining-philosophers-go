// ./philo numPhilos timeDie timeEat timeSleep [mealsRequired]
package main

import (
    "strconv"
    "time"
    "fmt"
)

type Program struct {
	
	numPhilos int

	timeDie time.Duration
	timeEat time.Duration
	timeSleep time.Duration
	
	mealsRequired int

	//ctx context.Context // this is an interface, a shared signal says "program must stop now"
}

//this function assigns each args to the equivalent in the struct
func initProgram(program Program) {
	
}