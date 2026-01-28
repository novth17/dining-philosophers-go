package main //all files in this folder belong to the same exe program

import (
	"fmt"
	"strconv"
)

func validateArgs(args []string) error {
	
	if len(args) != 4 && len(args) != 5 {
        return fmt.Errorf("usage: ./philo numPhilos timeDie timeEat timeSleep [mealsRequired]")
	}

	for i := 0; i < len(args); i++ {
		_, err := parsePositiveNumber(args[i])
		if err != nil {
			return err
		}
	}
	return nil
}

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
