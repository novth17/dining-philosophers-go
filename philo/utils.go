package main

import (
	"fmt"
	"os"
	"runtime"
	"sync/atomic"
	"time"
)

//improvement: considering logger channel, lock and unlock printing can cause some contention. how oftern you are printing
func (program *Program) printMtx(id int, action string) bool {
    // Fast-path: check stopSim WITHOUT locking
    if atomic.LoadInt32(&program.stopSim) == 1 {
        return false
    }

    program.logMu.Lock()
    defer program.logMu.Unlock()

    // Re-check after locking to prevent race conditions
    if program.stopSim == 1 {
        return false
    }

    timestamp := time.Since(program.startTime).Milliseconds()
    fmt.Fprintf(os.Stdout, "%d %d %s\n", timestamp, id, action)

    // If this log is a death, we stop the sim while holding the lock
    if action == "died" {
        atomic.StoreInt32(&program.stopSim, 1)
    }
    return true
}

func (program *Program) printCPUInfo() {
    fmt.Printf("Logical CPUs (Hardware): %d\n", runtime.NumCPU())
    fmt.Printf("Go Scheduler Parallelism (GOMAXPROCS): %d\n", runtime.GOMAXPROCS(0))
}

func msToDuration(ms int) time.Duration {
	return time.Duration(ms) * time.Millisecond
}

func (p *Program) Stop() {
	if p.cancel != nil {
		p.cancel()
	}
}
