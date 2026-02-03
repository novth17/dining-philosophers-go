package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func (program *Program) Run() {
    fmt.Println("--- Program Running.... ---")

    mux := http.NewServeMux()
    mux.Handle("/", http.FileServer(http.Dir("./web")))
    mux.HandleFunc("/events", program.eventsHandler)

    // 1. Start the server in a goroutine so it doesn't block this function
    go func() {
        if err := http.ListenAndServe(":8080", mux); err != nil {
            fmt.Printf("Server failed: %v\n", err)
        }
    }()

    fmt.Println("Waiting for browser connection at http://localhost:8080...")

    // 2. NOW wait here. This will block until program.StartSignal is closed
    // inside your eventsHandler function.
    <-program.StartSignal

    // 3. Only now do we start the clock and the routines
    now := time.Now()
    program.startTime = now

    for i := 0; i < program.numPhilos; i++ {
        program.philos[i].timeLastMeal = now
    }
    
    program.emit(Event{
        Time:  0,
        Event: "simulation_start",
    })

    ctx, cancel := context.WithCancel(context.Background()) // empty context
    defer cancel()

	go program.monitor(ctx, cancel)

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
