package main

import (
	"sync"
	//"time"
)

func initProgram(program *Program, args []string) error {

	var err	error
	program.numPhilos, err = parseValidNumber(args[0])
	if err != nil {
		return err
	}

	timeDieMs, err := parseValidNumber(args[1])
	if err != nil {
		return err
	}
	program.timeDie = msToDuration(timeDieMs)

	timeEatMs, err := parseValidNumber(args[2])
	if err != nil {
		return err
	}
	program.timeEat = msToDuration(timeEatMs)

	timeSleepMs, err := parseValidNumber(args[3])
	if err != nil {
		return err
	}
	program.timeSleep = msToDuration(timeSleepMs)

	if len(args) == 5 {
		program.mealsRequired, err = parseValidNumber(args[4])
		if err != nil {
			return err
		}
	} else {
		program.mealsRequired = -1
	}
	//program.startTime = time.Now() - start in run
	program.Events = make(chan Event, 100)
	program.forks = make([]sync.Mutex, program.numPhilos)
	program.StartSignal = make(chan struct{})

	if err := initPhilo(program); err != nil {
		return err
	}
	return nil
}

func initPhilo(program *Program) error {

	program.philos = make([]Philo, program.numPhilos)

	for i := 0; i < program.numPhilos; i++ {
		program.philos[i] = Philo{
			id: i + 1,
			prog: program,
			//timeLastMeal: time.Now(),
			mealCount: 0,
		}
	}
	return nil
}
