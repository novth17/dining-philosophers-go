package main //all files in this folder belong to the same exe program

import (
	"os"
	"fmt"
)

func main() {

	//create the args
	executionArg := os.Args
	argsWithoutExe := os.Args[1:] // from index 1 until the end

    fmt.Println(executionArg)
    fmt.Println(argsWithoutExe)

	// validation
	err:= validateArgs(argsWithoutExe)
	if err != nil {
		fatal(err)
	}

	//program creation, this is where args are being assigned
	program, err:= NewProgram(argsWithoutExe)
	if err != nil {
		fatal(err)
	}

	//program start
	program.Run()
}