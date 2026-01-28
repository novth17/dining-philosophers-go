package main

import (
	"context"
	"fmt"
	"time"
)

func (program *Program) monitor(ctx context.Context, cancel context.CancelFunc) {
    for {
        // Stop monitoring if simulation already ended (ctx.Done() or cancel() called elsewhere)
        if ctx.Err() != nil {
            return
        }

        for i := 0; i < program.numPhilos; i++ {
            philo := &program.philos[i]

			if philo.isStarving() {
				program.logMutex.Lock() 
				program.stopSim = true
				ms := time.Since(program.startTime).Milliseconds()
				fmt.Printf("%d %d died\n", ms, philo.id)
				program.logMutex.Unlock()
				
				cancel() 
				return
			}
		}

        if program.mealsRequired != -1 && program.checkAllFull() {
            cancel()
            return
        }
        time.Sleep(500 * time.Microsecond)
    }
}

func (philo *Philo) isStarving() bool {
	philo.prog.mealMutex.Lock()
	defer philo.prog.mealMutex.Unlock()
	return time.Since(philo.timeLastMeal) > philo.prog.timeDie
}

func (program *Program) checkAllFull() bool {
    program.mealMutex.Lock()
    defer program.mealMutex.Unlock()

    fullPhilos := 0
    for i := 0; i < program.numPhilos; i++ {
        if program.philos[i].mealCount >= program.mealsRequired {
            fullPhilos++
        }
    }
    return fullPhilos == program.numPhilos
}