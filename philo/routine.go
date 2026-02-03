package main

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
)
func (philo *Philo) routine(ctx context.Context) {
    // Initial delay for even IDs to prevent immediate deadlock
    if philo.id%2 == 0 {
        time.Sleep(time.Millisecond)
    }

    for {
        if !philo.eat(ctx) { return }
        if !philo.sleep(ctx) { return }
        if !philo.think(ctx) { return }
    }
}

func (philo *Philo) think(ctx context.Context) bool {
	atomic.StoreInt32(&philo.State, THINKING)
	if !philo.prog.printMtx(philo.id, "is thinking") {
		return false
	}


	if philo.prog.numPhilos%2 != 0 {
		// Formula: (time_to_eat * 2) - time_to_sleep. This ensures we wait long enough for the other pair to finish.
		thinkTime := (philo.prog.timeEat * 2) - philo.prog.timeSleep
		
		// But don't think so long that YOU die. 
		// Usually, a small 10-20ms gap is enough to fix the race.
		if thinkTime < 0 {
			thinkTime = 10 * time.Millisecond
		} else {
			thinkTime = thinkTime / 2 // Start with half the gap
		}
		
		return philo.safeSleep(thinkTime, ctx)
	}

	return true
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
	if ctx.Err() != nil {
		second.Unlock()
		first.Unlock()
		return false
	}

	program.printMtx(philo.id, "has taken a fork")

	program.mealMu.Lock()
	philo.timeLastMeal = time.Now()
	philo.mealCount++
	program.mealMu.Unlock()

	atomic.StoreInt32(&philo.State, EATING)
	program.printMtx(philo.id, "is eating")
	eatSuccess := philo.safeSleep(program.timeEat, ctx)

	second.Unlock()
	first.Unlock()
	return eatSuccess
}

func (philo *Philo) sleep(ctx context.Context) bool {
	atomic.StoreInt32(&philo.State, SLEEPING)
	philo.prog.printMtx(philo.id, "is sleeping")
	return philo.safeSleep(philo.prog.timeSleep, ctx)
}


func (philo *Philo) safeSleep(duration time.Duration, ctx context.Context) bool {
	start := time.Now()
	for time.Since(start) < duration {
		if ctx.Err() != nil {
			return false
		}
		time.Sleep(time.Millisecond)
	}
	return true
}