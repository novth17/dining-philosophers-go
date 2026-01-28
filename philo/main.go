package main //all files in this folder belong to the same exe program
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
	program.Run()
	program.printCPUInfo()
}