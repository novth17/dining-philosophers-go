package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

func (program *Program) monitor(ctx context.Context) {
    for {
        if ctx.Err() != nil {
            return
        }

        for i := 0; i < program.numPhilos; i++ {
            if program.philos[i].isStarving() {
            program.logMu.Lock()
            // Check if someone else already ended the party
            //The atomic package uses special CPU instructions (like LOCK in x86 or LDAXR/STLXR in ARM/M-series) that lock that specific memory address for a few nanoseconds. Prevents "Dirty Reads" & ensure Visibility
            if atomic.LoadInt32(&program.stopSim) == 0 {
                atomic.StoreInt32(&program.stopSim, 1)
                fmt.Printf("%d %d died\n", time.Since(program.startTime).Milliseconds(), program.philos[i].id)
                program.cancel()
            }
            program.logMu.Unlock()
            return
            }
        }

        if program.mealsRequired != -1 && program.checkAllFull() {
            program.cancel()
            return
        }
        
        //monitor wakes immediately on cancellation
        select {
        case <-ctx.Done():
            return
        case <-time.After(500 * time.Microsecond):
        }
    }
}

func (philo *Philo) isStarving() bool {
    philo.prog.mealMu.Lock()
    defer philo.prog.mealMu.Unlock()
    return time.Since(philo.timeLastMeal) >= philo.prog.timeDie // improvement if production: can there be a buffer? small extra window based on real world behavior
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