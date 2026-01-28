package main

import (
	"fmt"
	"os"
	"time"
	"runtime"
)

func (philo *Philo) isStarving() bool {
	philo.prog.mealMutex.Lock()
	defer philo.prog.mealMutex.Unlock()
	return time.Since(philo.timeLastMeal) > philo.prog.timeDie
}

func msToDuration(ms int) time.Duration {
	return time.Duration(ms) * time.Millisecond
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, "Error: ", err)
	os.Exit(1)
}

func (program *Program) printMtx(args ...any) bool {
	ms := time.Since(program.startTime).Milliseconds()

	program.logMutex.Lock()
	defer program.logMutex.Unlock()

	if program.stopSim {
		return false
	}
	fmt.Fprintln(os.Stdout, append([]any{ms}, args...)...)
	return true
}


func (program *Program) printCPUInfo() {
    fmt.Printf("Logical CPUs (Hardware): %d\n", runtime.NumCPU())
    fmt.Printf("Go Scheduler Parallelism (GOMAXPROCS): %d\n", runtime.GOMAXPROCS(0))
}

//learn to love defer: if err != nil {return} --> never unlock
