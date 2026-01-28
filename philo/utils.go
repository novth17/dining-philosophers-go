package main

import (
	"fmt"
	"os"
	"time"
	"runtime"
)

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
