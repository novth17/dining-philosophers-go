// ./philo numPhilos timeDie timeEat timeSleep [mealsRequired]
package main

import (
	"time"
)

type Program struct {
	
	numPhilos int

	timeDie time.Duration
	timeEat time.Duration
	timeSleep time.Duration
	
	mealsRequired int

	//ctx context.Context // this is an interface, a shared signal says "program must stop now"
}

func msToDuration(ms int) time.Duration {
	return time.Duration(ms) * time.Millisecond
}

//this function assigns each args to the equivalent in the struct
func NewProgram(args []string) (*Program, error) {

	var program Program
	var err		error

	program.numPhilos, err = parsePositiveNumber(args[0])
	if err != nil {
		return nil, err
	}
	timeDieMs, err := parsePositiveNumber(args[1])
	if err != nil {
		return nil, err
	}
	program.timeDie = msToDuration(timeDieMs)

	timeEatMs, err := parsePositiveNumber(args[2])
	if err != nil {
		return nil, err
	}
	program.timeEat = msToDuration(timeEatMs)

	timeSleepMs, err := parsePositiveNumber(args[3])
	if err != nil {
		return nil, err
	}
	program.timeSleep = msToDuration(timeSleepMs)

	if len(args) == 5 {
		program.mealsRequired, err = parsePositiveNumber(args[4])
		if err != nil {
			return nil, err
		}
	} else {
		program.mealsRequired = -1 //-1 can never be a valid meal count as only pos accepted
	}

	return &program, nil
}