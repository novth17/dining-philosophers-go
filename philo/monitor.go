package main

import (
	"context"
	"time"
)

func (program *Program) monitor(ctx context.Context, cancel context.CancelFunc) {
    for {
        if ctx.Err() != nil {
            return
        }

        for i := 0; i < program.numPhilos; i++ {
            if program.philos[i].isStarving() {
                // 1. SIGNAL IMMEDIATELY
                cancel() 
                
                // 2. LOG IT (Log internal logic will set stopSim = 1)
                program.Log(program.philos[i].id, "died")
                return
            }
        }

        if program.mealsRequired != -1 && program.checkAllFull() {
            cancel()
            return
        }
        // Tightening this to 100-500us is good for 10ms timers
        time.Sleep(500 * time.Microsecond)
    }
}

func (philo *Philo) isStarving() bool {
    philo.prog.mealMu.Lock()
    defer philo.prog.mealMu.Unlock()
    
    // Use >= to be strict on the death timer
    return time.Since(philo.timeLastMeal) >= philo.prog.timeDie
}

func (program *Program) checkAllFull() bool {
    program.mealMu.Lock()
    defer program.mealMu.Unlock()

    fullPhilos := 0
    for i := 0; i < program.numPhilos; i++ {
        if program.philos[i].mealCount >= program.mealsRequired {
            fullPhilos++
        }
    }
    return fullPhilos == program.numPhilos
}