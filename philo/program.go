package main

import (
	"sync"
	"time"
	"context"
)
type Philo struct {
	id 				int
	prog 			*Program
	timeLastMeal 	time.Time
	mealCount 		int
}

type Program struct {
	numPhilos		int
	timeDie			time.Duration
	timeEat			time.Duration
	timeSleep		time.Duration
	mealsRequired 	int 
	startTime 		time.Time
	stopSim 		int32
	// 	slice header
	//  ├─ data *T ─────► [ Philo ][ Philo ][ Philo ]
	//  ├─ len = 3
	//  └─ cap = 3
	//s := make([]int, 3, 10)
	philos			[]Philo
	forks			[]sync.Mutex

	logMu			sync.Mutex
	mealMu			sync.Mutex
	wg    			sync.WaitGroup
	ctx				context.Context
	cancel			context.CancelFunc
}

func NewProgram(args []string) (*Program, error) {
	var program Program

	ctx, cancel := context.WithCancel(context.Background())

	program.ctx = ctx
	program.cancel = cancel

	err := initProgram(&program, args)
	if err != nil {
		return nil, err
	}

	return &program, nil
}
