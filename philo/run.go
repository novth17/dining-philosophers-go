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
}