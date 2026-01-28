package main

import (
	"context"
	"sync"
	"time"
)

func (philo *Philo) isStarving() bool {
	philo.prog.mealMutex.Lock()
	defer philo.prog.mealMutex.Unlock()
	return time.Since(philo.timeLastMeal) > philo.prog.timeDie
}

func (philo *Philo) routine(ctx context.Context) {  
	program := philo.prog

	if (philo.id % 2 == 0) {
		philo.think(ctx);
		philo.safeSleep(program.timeEat / 2, ctx);
	}

	for {
		if !philo.eat(ctx) { return }
		if !philo.sleep(ctx) { return }
		if !philo.think(ctx) { return }
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

	// --- acquire forks ---
	first.Lock()
	program.printMtx(philo.id, "has taken a fork")

	// single philosopher case
	if program.numPhilos == 1 {
		first.Unlock()
		select {
		case <-time.After(program.timeDie + 10*time.Millisecond):
			return false
		case <-ctx.Done():
			return false
		}
	}

	second.Lock()
	program.printMtx(philo.id, "has taken a fork")

	// update meal info
	program.mealMutex.Lock()
	philo.timeLastMeal = time.Now()
	philo.mealCount++
	program.mealMutex.Unlock()

	program.printMtx(philo.id, "is eating")

	// --- ACTUALLY EAT (forks still held) ---
	eatSuccess := philo.safeSleep(program.timeEat, ctx)

	// --- release forks immediately after eating ---
	second.Unlock()
	first.Unlock()

	return eatSuccess
}


//sleep
func (philo *Philo) sleep(ctx context.Context) bool {
	program := philo.prog
	program.printMtx(philo.id, "is sleeping")
	return philo.safeSleep(program.timeSleep, ctx)
}

//think
func (philo *Philo) think(ctx context.Context) bool {
	program := philo.prog
	program.printMtx(philo.id, "is thinking")

	// 1. For even numbers (2, 4, 6...), we don't need a math formula.
	// Just yield to the scheduler for a tiny bit (or 0) to keep things fast.
	if program.numPhilos%2 == 0 {
		return philo.safeSleep(time.Millisecond, ctx)
	}

	// 2. For odd numbers (5, 7, 9...), we calculate a "Rotation Delay".
	// The ideal time is (Eat * 2) - Sleep.
	// This lets your neighbors finish their turn without you starving.
	thinkTime := (program.timeEat * 2) - program.timeSleep
	if thinkTime < 0 {
		thinkTime = 0
	}

	// 3. THE SAFETY CAP: Never think for more than 20% of your remaining sad life.
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
