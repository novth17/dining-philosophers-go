package main

import (
	"sync"
	"time"
)

func initProgram(program *Program, args []string) error {

	var err	error

	program.numPhilos, err = parsePositiveNumber(args[0])
	if err != nil {
		return err
	}

	timeDieMs, err := parsePositiveNumber(args[1])
	if err != nil {
		return err
	}
	program.timeDie = msToDuration(timeDieMs)

	timeEatMs, err := parsePositiveNumber(args[2])
	if err != nil {
		return err
	}
	program.timeEat = msToDuration(timeEatMs)

	timeSleepMs, err := parsePositiveNumber(args[3])
	if err != nil {
		return err
	}
	program.timeSleep = msToDuration(timeSleepMs)

	if len(args) == 5 {
		program.mealsRequired, err = parsePositiveNumber(args[4])
		if err != nil {
			return err
		}
	} else {
		program.mealsRequired = -1
	}
	program.startTime = time.Now()
	program.forks = make([]sync.Mutex, program.numPhilos)

	if err := initPhilo(program); err != nil {
		return err
	}
	return nil
}

func initPhilo(program *Program) error {

	//Slices can be created with the built-in make
	program.philos = make([]Philo, program.numPhilos)

	for i := 0; i < program.numPhilos; i++ {
		program.philos[i] = Philo{
			id: i + 1,
			prog: program,
			timeLastMeal: time.Now(),
			mealCount: 0,
		}
	}
	return nil
}
