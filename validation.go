package main //all files in this folder belong to the same exe program

import (
	"fmt"
	"strconv"
)

func parsePositiveNumber(str string) (int, error) {
    number, err := strconv.Atoi(str)
    if err != nil {
        return 0, fmt.Errorf("not a number: %s", str)
    }
    if number <= 0 {
        return 0, fmt.Errorf("number must be positive: %s", str)
    }
    return number, nil
}

func validateArgs(args []string) error {

	fmt.Println("--- Validating args...---")

	if len(args) != 4 && len(args) != 5 {
        return fmt.Errorf("usage: ./philo numPhilos timeDie timeEat timeSleep [mealsRequired]")
	}

	for i := 0; i < len(args); i++ {
		_, err := parsePositiveNumber(args[i]) //_ means I intentionally ignore this var
		if err != nil {
			return err
		}
	}
	return nil
}