package main

import (
	"fmt"
	"os"
	"runtime"
	"sync/atomic"
	"time"
)

func msToDuration(ms int) time.Duration {
	return time.Duration(ms) * time.Millisecond
}
func (program *Program) printMtx(args ...any) bool {
    if atomic.LoadInt32(&program.stopSim) == 1 {
        return false
    }

    ms := time.Since(program.startTime).Milliseconds()

    program.logMu.Lock()
    defer program.logMu.Unlock()

    if program.stopSim == 1 {
        return false
    }
    
    fmt.Fprintln(os.Stdout, append([]any{ms}, args...)...)
    return true
}

func (program *Program) printCPUInfo() {
    fmt.Printf("Logical CPUs (Hardware): %d\n", runtime.NumCPU())
    fmt.Printf("Go Scheduler Parallelism (GOMAXPROCS): %d\n", runtime.GOMAXPROCS(0))
}

