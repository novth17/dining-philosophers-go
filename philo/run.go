package main

import (
	"context"
	"fmt"
)

//method with a pointer receiver on Program.

func (program *Program) Run() {
    fmt.Println("--- Pupu The Host brought a big bowl of ice cream. The party is on! ---")

    // 1. Setup the Megaphone and the Red Button
    ctx, cancel := context.WithCancel(context.Background()) // empty context
    defer cancel()

	//start monitor goroutine
	go program.monitor(ctx, cancel)

    // Start Philosophers
    for i := 0; i < program.numPhilos; i++ {
        program.wg.Add(1)
		 fmt.Println("--- start philosopher! ---")
        go func(p *Philo) {
            defer program.wg.Done()
            p.routine(ctx)
        }(&program.philos[i])
    }

    // Wait for everyone to respect the cancel() signal
    program.wg.Wait()
    fmt.Println("--- Party's over! ---")

    // Check if mealsRequired was actually set (not -1)
    if program.mealsRequired != -1 {
        program.mealMutex.Lock()
        defer program.mealMutex.Unlock()

        fmt.Println("\n--- Final Meal Statistics ---")
        for i := 0; i < program.numPhilos; i++ {
            fmt.Printf("Philosopher %d finished %d/%d meals\n", 
                program.philos[i].id, 
                program.philos[i].mealCount, 
                program.mealsRequired)
        }
    }
}
