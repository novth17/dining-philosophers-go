// ./philo numPhilos timeDie timeEat timeSleep [mealsRequired]
package main

import (
	"sync"
	"time"
)
type Philo struct {
	id int
	prog *Program
	timeLastMeal time.Time
	mealCount int //need mutex or atomic
}

type Program struct {
	numPhilos int
	timeDie time.Duration
	timeEat time.Duration
	timeSleep time.Duration
	mealsRequired int 
	startTime time.Time
	stopSim int32

	// 	slice header
	//  ├─ data ─────► [ Philo ][ Philo ][ Philo ]
	//  ├─ len = 3
	//  └─ cap = 3
	philos []Philo //slice of struct Philo
	forks []sync.Mutex //slice of struct of fork mutexes

	logs      []string
	logMu sync.Mutex
	mealMu sync.Mutex
	wg         sync.WaitGroup

	//for frontend
	Events	chan Event

	StartSignal chan struct{} // closed when the first client connects
    once        sync.Once //ensures we only close the channel once
}

//this function assigns each args to the equivalent in the struct
func NewProgram(args []string) (*Program, error) {

	var program Program
	
	err := initProgram(&program, args)
	if (err != nil) {
		return nil, err
	}
	return &program, nil
}

func (program *Program) Now() int64 {
	return time.Since(program.startTime).Milliseconds()
}