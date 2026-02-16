package main

import (
	"context"
	"fmt"
)

func (program *Program) Run() {
    fmt.Println("--- Program Running.... ---")

    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    //add wg for monitor, should wait everything 
    program.wg.Add(1)
    go func() {
        defer program.wg.Done()
        program.monitor(ctx, cancel)
    }()

    for i := 0; i < program.numPhilos; i++ {
        program.wg.Add(1)
        go func(p *Philo) {
            defer program.wg.Done()
            p.routine(ctx)
        }(&program.philos[i])
    }

    program.wg.Wait()
    fmt.Println("--- Party's over! ---")

    if program.mealsRequired != -1 {
        program.mealMu.Lock()
        defer program.mealMu.Unlock()

        fmt.Println("\n--- Final Meal Statistics ---")
        for i := 0; i < program.numPhilos; i++ {
            fmt.Printf("Philosopher %d finished %d/%d meals\n", 
                program.philos[i].id, 
                program.philos[i].mealCount, 
                program.mealsRequired)
        }
    }
}
