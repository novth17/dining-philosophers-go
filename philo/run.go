package main

import (
	"fmt"
)

//method with a pointer receiver on Program.
func (program *Program) Run() {
	//fmt.Printf("Program initialized: %+v\n", program)
	fmt.Println("--- Pupu The Host brought a big bowl of ice cream. The party is on! ---")

	//start one coroutine per one philo
	for i := 0; i < program.numPhilos; i++ {
		go program.philos[i].routine()
	}

	select {}
	//wait until sim ends
}