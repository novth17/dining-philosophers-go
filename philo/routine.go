package main

import (
	"context"
	"sync"
	"time"
)

func (philo *Philo) routine(ctx context.Context) {  
	program := philo.prog

	if (philo.id % 2 == 0) {
		philo.think(ctx);
		philo.safeSleep(program.timeEat / 2, ctx);
	}

	for {
		if !philo.eat(ctx) || !philo.sleep(ctx) || !philo.think(ctx) {
			return
		}
	}
}

func (philo *Philo) eat(ctx context.Context) bool {
	program := philo.prog
	left := philo.id - 1
	right := philo.id % program.numPhilos

	var first, second *sync.Mutex
	if philo.id%2 == 0 {
		first, second = &program.forks[left], &program.forks[right]
	} else {
		first, second = &program.forks[right], &program.forks[left]
	}

	first.Lock()
	program.printMtx(philo.id, "has taken a fork")

	if program.numPhilos == 1 {
		first.Unlock()
		<-ctx.Done() 
		return false
	}

	second.Lock()

    // Did we die while waiting for this lock?
    if ctx.Err() != nil {
        second.Unlock()
        first.Unlock()
        return false
    }

	program.printMtx(philo.id, "has taken a fork")

	program.mealMutex.Lock()
	philo.timeLastMeal = time.Now()
	philo.mealCount++
	program.mealMutex.Unlock()

	program.printMtx(philo.id, "is eating")

	eatSuccess := philo.safeSleep(program.timeEat, ctx)

	second.Unlock()
	first.Unlock()

	return eatSuccess
}

func (philo *Philo) sleep(ctx context.Context) bool {
	program := philo.prog
	program.printMtx(philo.id, "is sleeping")
	return philo.safeSleep(program.timeSleep, ctx)
}

func (philo *Philo) think(ctx context.Context) bool {
	program := philo.prog
	program.printMtx(philo.id, "is thinking")

	if program.numPhilos%2 == 0 {
		return philo.safeSleep(time.Millisecond, ctx)
	}

	// For odd numbers, calculate a "Rotation Delay" to avoid stealing 
	thinkTime := (program.timeEat * 2) - program.timeSleep
	if thinkTime < 0 {
		thinkTime = 0
	}

	// Never think for more than 20% of remaining sad life.
	if thinkTime > program.timeDie / 5 {
		thinkTime = program.timeDie / 5
	}

	return philo.safeSleep(thinkTime, ctx)
}

func (philo *Philo) safeSleep(dur time.Duration, ctx context.Context) bool {
	select {
	case <-time.After(dur):
		return true
	case <-ctx.Done():
		return false
	}
}
