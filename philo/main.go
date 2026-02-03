package main

import (
	"os"
)

func main() {

	argsWithoutExe := os.Args[1:]

	err:= validateArgs(argsWithoutExe)
	if err != nil {
		fatal(err)
	}
	program, err:= NewProgram(argsWithoutExe)
	if err != nil {
		fatal(err)
	}
	program.Run(":8080")
	program.printCPUInfo()
}