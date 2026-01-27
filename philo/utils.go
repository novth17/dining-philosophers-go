package main

import (
	"fmt"
	"os"
	"time"
)

func msToDuration(ms int) time.Duration {
	return time.Duration(ms) * time.Millisecond
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, "Error: ", err)
	os.Exit(1)
}

func printBasic(args ...any) {
	fmt.Fprintln(os.Stdout, args...)
}

func (program *Program) printMtx(args ...any) {
	
	ms := time.Since(program.startTime).Milliseconds()
	program.logMutex.Lock()

	//defer: run this line when the function returns, not now
	defer program.logMutex.Unlock()
	fmt.Fprintln(os.Stdout, append([]any{ms}, args...)...)
}

//learn to love defer: if err != nil {return} --> never unlock
