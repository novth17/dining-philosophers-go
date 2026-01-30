package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func (p *Program) Log(id int, action string) {

	if atomic.LoadInt32(&p.stopSim) == 1 {
        return
    }
	p.logMu.Lock()
    defer p.logMu.Unlock()
    
    if p.stopSim == 1 { 
        return 
    }

    timestamp := time.Since(p.startTime).Milliseconds()
    msg := fmt.Sprintf("%d %d %s", timestamp, id, action)
    p.logs = append(p.logs, msg)
    
    if action == "died" { 
        p.stopSim = 1 
    }
}

func (p *Program) FinalPrint() {
    for _, line := range p.logs {
        fmt.Println(line)
    }
}