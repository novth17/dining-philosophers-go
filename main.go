package main //all files in this folder belong to the same exe program

import (
	"os"
	"fmt"
)

func main() {

	//create the args

	var executionArg []string = os.Args
	var argsWithoutExe = os.Args[1:] // from index 1 until the end

    fmt.Println(executionArg)
    fmt.Println(argsWithoutExe)

	// validation
	err:= validateArgs(argsWithoutExe)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	//program creation
	

	//program start
	
	fmt.Println("--- Philo Pupu opened the shop ---")

}