package main

import (
	"context"
	"fmt"
	"time"
)

func (program *Program) monitor(ctx context.Context, cancel context.CancelFunc) {
	for {
		if ctx.Err() != nil {
			return
		}

		for i := 0; i < program.numPhilos; i++ {
			if program.philos[i].isStarving() {
				program.logMu.Lock()
				program.stopSim = 1

				now := time.Since(program.startTime).Milliseconds()

				//Death event initiated
				program.emit(Event{
					Time:  now,
					Philo: program.philos[i].id,
					Event: "death",
				})

				fmt.Printf("%d %d died\n", now, program.philos[i].id)

				cancel()
				program.logMu.Unlock()
				return
			}
		}

		if program.mealsRequired != -1 && program.checkAllFull() {
			now := time.Since(program.startTime).Milliseconds()

			// Simulation end event (all full)
			program.emit(Event{
				Time:  now,
				Event: "simulation_end",
			})

			cancel()
			return
		}

		time.Sleep(500 * time.Microsecond)
	}
}

func (philo *Philo) isStarving() bool {
    program := philo.prog

	philo.prog.mealMu.Lock()
	defer philo.prog.mealMu.Unlock()

    if time.Since(philo.timeLastMeal) > program.timeDie*8/10 {
    program.emit(Event{
        Time:  program.Now(),
        Philo: philo.id,
        State: "hungry",
    })
}

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

