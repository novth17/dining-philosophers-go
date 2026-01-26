package main //all files in this folder belong to the same exe program

import (
	"fmt"
	"strconv"
)

func parsePositiveNumber(str string) (int, error) {
    var number int
    var err error

    number, err = strconv.Atoi(str)
    if err != nil {
        return 0, fmt.Errorf("Not a number: %s", str)
    }
	if (number <= 0) {
		return 0, fmt.Errorf("Number must be positive: %s", str)
	}
    return number, nil
}


func validateArgs(args []string) error {

	fmt.Println("--- Validating args...---")

	if (len(args) != 4 && len(args) != 5) {
		return fmt.Errorf("Not correct arguments. Usage: ./philo numPhilos timeDie timeEat timeSleep [mealsRequired]")
	}

	for i := 0; i < len(args); i++ {
		number, error := parsePositiveNumber(args[i])
		if error != nil {
			return error
		}
		_ = number // means I intentionally ignore this var
	}

	return nil
}